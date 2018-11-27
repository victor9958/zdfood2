/**

 @Name：layuiAdmin 主入口
 @Author：贤心
 @Site：http://www.layui.com/admin/
 @License：LPPL

 */

layui.extend({
    setter: 'config' //配置文件
    , admin: 'lib/admin' //核心模块
    , view: 'lib/view' //核心模块
}).define(['setter', 'admin'], function (exports) {
    var setter = layui.setter
        , element = layui.element
        , admin = layui.admin
        , tabsPage = admin.tabsPage
        , view = layui.view

        //根据路由渲染页面
        , renderPage = function () {
            var router = layui.router()
                , path = router.path
                , pathURL = admin.correctRouter(router.path.join('/'));


            //默认读取主页
            if (!path.length) path = [''];

            //如果最后一项为空字符，则读取默认文件
            if (path[path.length - 1] === '') {
                path[path.length - 1] = setter.entry;
            }

            /*
            layui.config({
              base: setter.base + 'controller/'
            });
            */

            //重置状态
            var reset = function (type) {
                renderPage.haveInit && layer.closeAll();
                renderPage.haveInit = true;
                $(APP_BODY).scrollTop(0);

                delete tabsPage.type; //重置页面标签的来源类型
            };

            //如果路由来自于 tab 切换，则不重新请求视图
            if (tabsPage.type === 'tab') {
                //切换到非主页、或者切换到主页且主页必须有内容。方可阻止请求
                if (pathURL !== '/' || (pathURL === '/' && admin.tabsBody().html())) {
                    admin.tabsBodyChange(tabsPage.index);
                    return reset(tabsPage.type);
                }
            }

            //请求视图渲染
            view().render(path.join('/')).then(function (res) {

                //遍历页签选项卡
                var matchTo
                    , tabs = $('#LAY_app_tabsheader>li');

                tabs.each(function (index) {
                    var li = $(this)
                        , layid = li.attr('lay-id');

                    if (layid === pathURL) {
                        matchTo = true;
                        tabsPage.index = index;
                    }
                });


                //如果未在选项卡中匹配到，则追加选项卡
                if (setter.pageTabs && pathURL !== '/') {
                    if (!matchTo) {
                        $(APP_BODY).append('<div class="layadmin-tabsbody-item layui-show"></div>');
                        tabsPage.index = tabs.length;
                        element.tabAdd(FILTER_TAB_TBAS, {
                            title: '<span>' + (res.title || '新标签页') + '</span>'
                            , id: pathURL
                            , attr: router.href
                        });
                    }
                }

                this.container = admin.tabsBody(tabsPage.index);

                //定位当前tabs
                element.tabChange(FILTER_TAB_TBAS, pathURL);
                admin.tabsBodyChange(tabsPage.index);

            }).done(function () {
                layui.use('common', layui.cache.callback.common);
                $win.on('resize', layui.data.resize);

                element.render('breadcrumb', 'breadcrumb');

                //容器 scroll 事件，剔除吸附层
                admin.tabsBody(tabsPage.index).on('scroll', function () {
                    var othis = $(this)
                        , elemDate = $('.layui-laydate')
                        , layerOpen = $('.layui-layer')[0];

                    //关闭 layDate
                    if (elemDate[0]) {
                        elemDate.each(function () {
                            var thisElemDate = $(this);
                            thisElemDate.hasClass('layui-laydate-static') || thisElemDate.remove();
                        });
                        othis.find('input').blur();
                    }

                    //关闭 Tips 层
                    layerOpen && layer.closeAll('tips');
                });
            });

            reset();
        }

        //入口页面
        , entryPage = function (fn) {
            var router = layui.router()
                , container = view(setter.container)
                , pathURL = admin.correctRouter(router.path.join('/'))
                , isIndPage;

            console.log(123)
            console.log(router)
            //检查是否属于独立页面
            layui.each(setter.indPage, function (index, item) {
                if (pathURL === item) {
                    return isIndPage = true;
                }
            });

            //将模块根路径设置为 controller 目录
            layui.config({
                base: setter.base + 'controller/'
            });

            //独立页面
            if (isIndPage || pathURL === '/auth/login') { //此处单独判断登入页，是为了兼容旧版（即未在 config.js 配置 indPage 的情况）
                container.render(router.path.join('/')).done(function () {
                    admin.pageType = 'alone';
                });
            } else { //后台框架页面

                //强制拦截未登入
                if (setter.interceptor) {
                    var local = layui.data(setter.tableName);
                    if (!local[setter.request.tokenName]) {
                        return location.hash = '/auth/login/redirect=' + encodeURIComponent(pathURL); //跳转到登入页
                    }
                }

                //渲染后台结构
                if (admin.pageType === 'console') { //后台主体页
                    renderPage();
                } else { //初始控制台结构
                    if (router.path[0] === '') {
                        // layui.data('menuParam', {
                        //   key: 'id',
                        //   value: 1
                        // });
                    } else if(router.path.length === 0) {
                        location.hash = '/';
                    }else {
                        layui.each(layui.data('baseMenu').menus, function (index, item) {
                            if (router.path[0].indexOf(item.name) != -1) {
                                layui.data('menuParam', {key: 'id', value: item.id});
                            }
                        });
                    }

                    container.render('layout').done(function () {
                        renderPage();
                        layui.element.render();

                        if (admin.screen() < 2) {
                            admin.sideFlexible();
                        }
                        admin.pageType = 'console';
                    });
                }
                layui.data('temporary', {key: 'queryString', value: ''})
            }


            var load;
            $.ajaxSetup({
                headers: {'Authorization': layui.data(setter.tableName)[setter.request.tokenName]}
                , timeout: 10000
                , beforeSend: function () {
                    load = layer.load(2, {
                        shade: [0.1, '#000'] //0.1透明度的白色背景
                    });
                    // view.loading;
                }
                , complete: function () {
                    layer.close(load)
                }
                , statusCode: {
                    401: function () {
                        view.exit();
                    },
                    400: function (e) {
                        layer.close(load);
                        layer.alert(JSON.parse(e.responseText).message, {icon: 5});
                    },
                    403: function (e) {
                        location.hash = '/template/tips/noauth';
                    },
                    404: function (e) {
                        location.hash = '/template/tips/404';
                    },
                    500: function () {
                        layer.close(load);
                        layer.alert('请求超时', {icon: 5, closeBtn: 0}, function () {
                            layui.index.render();
                        });
                    },
                    503: function () {
                        location.href = '/503';
                    }
                }
                , error: function (e, code) {
                    layer.close(load);
                    if (e.status == 401) {
                        view.exit();
                    } else if (e.status == 400) {
                        layer.alert(JSON.parse(e.responseText).message, {icon: 5});
                        // var error = [
                        //     '请求异常，请重试<br><cite>错误信息：</cite>'+code
                        //     ,debug()
                        // ].join('');
                        // view.error(error);
                        //
                        // typeof error === 'function' && error(res);
                    } else if (e.status == 403) {
                        location.hash = '/template/tips/noauth';
                    } else if (e.status == 404) {
                        location.hash = '/template/tips/404';
                    } else if (e.status == 503) {
                        location.href = '/503';
                    } else {
                        layer.alert('请求超时', {icon: 5, closeBtn: 0}, function () {
                            layui.index.render();
                        });
                    }
                }
            });
        }

        , APP_BODY = '#LAY_app_body', FILTER_TAB_TBAS = 'layadmin-layout-tabs'
        , $ = layui.$, $win = $(window);

    //初始主体结构
    layui.link(
        setter.base + 'style/admin2.css?v=' + (admin.v + '-1')
        , function () {
            entryPage()
        }
        , 'layuiAdmin'
    );

    layui.link(setter.base + 'style/admin.css?v=' + (admin.v + '-1'))
    layui.link(setter.base + 'style/common.css?v=' + (admin.v + '-1'))

    //监听Hash改变
    window.onhashchange = function () {
        entryPage()
        //执行 {setter.MOD_NAME}.hash 下的事件
        layui.event.call(this, setter.MOD_NAME, 'hash({*})', layui.router());


        var isIndPage = false, pathURL = admin.correctRouter(layui.router().path.join('/'));
        layui.each(setter.indPage, function (index, item) {
            if (pathURL === item) {
                return isIndPage = true;
            }
        });
        var path = layui.router().path.join('/'), url;
        if (path.substr(path.length - 1, 1) == '/') {
            path = path.substring(0, path.length - 1);
        }
        // console.log('dssfdsfs'+path);
        // console.log(layui.data('allMenu').menu.indexOf(path));
        if (!isIndPage && path == '') {
            url = '/';
            location.hash = url;
        } else if (!isIndPage && layui.data('allMenu').menu.indexOf(path) != -1 && layui.data('userMenu').menu.indexOf(path) == -1) {
            url = '/template/tips/noauth';
            location.hash = url;
        }
    };

    // $(document).on('click','#message',function(){
    //     $.each(layui.data('message'),function(i,k){
    //         if(k > 0){
    //             $('#'+i).append('<span class="layui-badge" style="right: 10px;">' + k + '</span>');
    //         }
    //     })
    // })

    //退出当前窗口时清空token
    /*window.onbeforeunload = function(){
        layui.data(setter.tableName, {
            key: setter.request.tokenName
            ,remove: true
        });
    };*/

    //扩展 lib 目录下的其它模块
    layui.each(setter.extend, function (index, item) {
        var mods = {};
        mods[item] = '{/}' + setter.base + 'lib/extend/' + item;
        layui.extend(mods);
    });

    //对外输出
    exports('index', {
        render: renderPage
    });
});
