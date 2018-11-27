/**

 @Name：layuiAdmin 核心模块
 @Author：贤心
 @Site：http://www.layui.com/admin/
 @License：LPPL
    
 */
 
layui.define('view', function(exports){
  var $ = layui.jquery
  ,laytpl = layui.laytpl
  ,element = layui.element
  ,setter = layui.setter
  ,view = layui.view
  ,device = layui.device()
  
  ,$win = $(window), $body = $('body')
  ,container = $('#'+ setter.container)
  
  ,SHOW = 'layui-show', HIDE = 'layui-hide', THIS = 'layui-this', TEMP = 'template'
  ,APP_BODY = '#LAY_app_body', APP_FLEXIBLE = 'LAY_app_flexible'
  ,FILTER_TAB_TBAS = 'layadmin-layout-tabs'
  ,APP_SPREAD_SM = 'layadmin-side-spread-sm', TABS_BODY = 'layadmin-tabsbody-item'
  ,ICON_SHRINK = 'layui-icon-shrink-right', ICON_SPREAD = 'layui-icon-spread-left'
  ,SIDE_SHRINK = 'layadmin-side-shrink', SIDE_MENU = 'LAY-system-side-menu'
      ,SIDE_MENU1 = 'LAY_menuItemElem', SIDE_MENU2 = 'LAY_menuContentElem'

  //通用方法
  ,admin = {
    v: '1.0.0-beta6'

    //数据的异步请求
    ,req: view.req

    //屏幕类型
    ,screen: function(){
      var width = $win.width()
      if(width >= 1200){
        return 3; //大屏幕
      } else if(width >= 992){
        return 2; //中屏幕
      } else if(width >= 768){
        return 1; //小屏幕
      } else {
        return 0; //超小屏幕
      }
    }

    //清除本地 token，并跳转到登入页
    ,exit: view.exit

    //侧边伸缩
    ,sideFlexible: function(status, tpl){
      var app = container
      ,iconElem =  $('#'+ APP_FLEXIBLE)
      ,screen = admin.screen();

      //如果没有二级菜单，则阻止操作
      if(!tpl && admin.iconFlexible('length')){
        // console.log("无三级菜单");
        return;
      }

      //设置状态，PC：默认展开、移动：默认收缩
      if(status === 'spread'){
        //切换到展开状态的 icon，箭头：←
        iconElem.removeClass(ICON_SPREAD).addClass(ICON_SHRINK);

        //移动：从左到右位移；PC：清除多余选择器恢复默认
        if(screen < 2){
          app.addClass(APP_SPREAD_SM);
            app.removeClass(SIDE_SHRINK);
        } else {
          app.removeClass(SIDE_SHRINK);
        }

        // app.removeClass(SIDE_SHRINK)
      } else {
        //切换到搜索状态的 icon，箭头：→
        iconElem.removeClass(ICON_SHRINK).addClass(ICON_SPREAD);

        //移动：清除多余选择器恢复默认；PC：从右往左收缩
        if(screen < 2){
          // app.removeClass(SIDE_SHRINK);
            app.addClass(APP_SPREAD_SM);
        }
          app.addClass(SIDE_SHRINK);

        // app.removeClass(APP_SPREAD_SM)
      }

      layui.event.call(this, setter.MOD_NAME, 'side({*})', {
        status: status
      });
    }

    //通过检查是否有二级菜单，才显示/隐藏"伸缩的icon"
    ,iconFlexible: function(length){
      var iconFlexible = $('.layadmin-flexible')
          ,menuShow = $('#'+ SIDE_MENU2).find('.layui-menu-item.layui-show')
          ,notMenu3 = menuShow.find('.layui-nav-item').length === 0;
      // console.log(menuShow[0], notMenu3);

      length ||iconFlexible[notMenu3 ? 'addClass' : 'removeClass'](HIDE);
      return notMenu3;
    }
    
    //事件监听
    ,on: function(events, callback){
      return layui.onevent.call(this, setter.MOD_NAME, events, callback);
    }
    
    //弹出面板
    ,popup: view.popup
    
    //右侧面板
    ,popupRight: function(options){
      //layer.close(admin.popup.index);
      return admin.popup.index = layer.open($.extend({
        type: 1
        ,id: 'LAY_adminPopupR'
        ,anim: -1
        ,title: false
        ,closeBtn: false
        ,offset: 'r'
        ,shade: 0.1
        ,shadeClose: true
        ,skin: 'layui-anim layui-anim-rl layui-layer-adminRight'
        ,area: '300px'
      }, options));
    }

    //中心面板
    ,popupCenter: function(options){
      //layer.close(admin.popup.index);
      return admin.popup.index = layer.open($.extend({
          title: false ,
          type: 1 ,
          resize: false,
          area: ['50%', '50%'] ,
          scrollbar: false ,
          offset: 'auto' ,
      }, options));
    }
    
    //主题设置
    ,theme: function(options){
      var theme = setter.theme
      ,local = layui.data(setter.tableName)
      ,id = 'LAY_layadmin_theme'
      ,style = document.createElement('style')

      ,styleText = laytpl([
        //主题色
        '.layui-side-menu,'
        ,'.layadmin-pagetabs .layui-tab-title li:after,'
        ,'.layadmin-pagetabs .layui-tab-title li.layui-this:after,'
        ,'.layui-layer-admin .layui-layer-title,'
        ,'.layadmin-side-shrink .layui-side-menu .layui-nav>.layui-nav-item>.layui-nav-child'
        ,'{background-color:{{d.color.main}} !important;}'
 
        //选中色
        ,'.layui-nav-tree .layui-this,'
        ,'.layui-nav-tree .layui-this>a,'
        ,'.layui-nav-tree .layui-nav-child dd.layui-this,'
        ,'.layui-nav-tree .layui-nav-child dd.layui-this a'
        ,'{background-color:{{d.color.selected}} !important;}'
        
        //logo
        ,'.layui-layout-admin .layui-logo{background-color:{{d.color.logo || d.color.main}} !important;}}'
      ].join('')).render(options = $.extend({}, local.theme, options))
      ,styleElem = document.getElementById(id);
      
      //添加主题样式
      if('styleSheet' in style){
        style.setAttribute('type', 'text/css');
        style.styleSheet.cssText = styleText;
      } else {
        style.innerHTML = styleText;
      }
      style.id = id;
      
      styleElem && $body[0].removeChild(styleElem);
      $body[0].appendChild(style);
      $body.attr('layadmin-themealias', options.color.alias);
      
      //本地存储记录
      local.theme = local.theme || {};
      layui.each(options, function(key, value){
        local.theme[key] = value;
      });
      layui.data(setter.tableName, {
        key: 'theme'
        ,value: local.theme
      }); 
    }
    
    //记录最近一次点击的页面标签数据
    ,tabsPage: {}
    
    //获取页面标签主体元素
    ,tabsBody: function(index){
      return $(APP_BODY).find('.'+ TABS_BODY).eq(index || 0);
    }
    
    //切换页面标签主体
    ,tabsBodyChange: function(index){
      admin.tabsBody(index).addClass(SHOW).siblings().removeClass(SHOW);
      events.rollPage('auto', index);
    }
    
    //resize事件管理
    ,resize: function(fn){
      var router = layui.router()
      ,key = router.path.join('-');
      $win.off('resize', admin.resizeFn[key]);
      fn(), admin.resizeFn[key] = fn;
      $win.on('resize', admin.resizeFn[key]);
    }
    ,resizeFn: {}
    ,runResize: function(){
      var router = layui.router()
      ,key = router.path.join('-');
      admin.resizeFn[key] && admin.resizeFn[key]();
    }
    ,delResize: function(){
      var router = layui.router()
      ,key = router.path.join('-');
      $win.off('resize', admin.resizeFn[key])
      delete admin.resizeFn[key];
    }
    
    //纠正路由格式
    ,correctRouter: function(href){
      if(!/^\//.test(href)) href = '/' + href;
      
      //纠正首尾
      return href.replace(/^(\/+)/, '/')
      .replace(new RegExp('\/' + setter.entry + '$'), '/'); //过滤路由最后的默认视图文件名（如：index）
    }
    
    //关闭当前 pageTabs
    ,closeThisTabs: function(){
      if(!admin.tabsPage.index) return;
      $(TABS_HEADER).eq(admin.tabsPage.index).find('.layui-tab-close').trigger('click');
    }
    
    //……
  };
  
  //事件
  var events = admin.events = {
    //伸缩
    flexible: function(othis){
      var iconElem = othis.find('#'+ APP_FLEXIBLE)
      ,isSpread = iconElem.hasClass(ICON_SPREAD);
      admin.sideFlexible(isSpread ? 'spread' : null);
    }
    
    //刷新
    ,refresh: function(){
      layui.index.render();
    }
    
    //点击消息
    ,message: function(othis){
      othis.find('.layui-badge-dot').remove();
      $.each(layui.data('message'),function(i,k){
          if(k > 0){
              $('#'+i).append('<span class="layui-badge" style="right: 10px;">' + k + '</span>');
          }
      })
    }
    
    //弹出主题面板
    ,theme: function(){
      admin.popupRight({
        id: 'LAY_adminPopupTheme'
        ,success: function(){
          view(this.id).render('system/theme', {
            id: 123
          })
        }
      });
    }
    
    //便签
    ,note: function(othis){
      var mobile = admin.screen() < 2
      ,note = layui.data(setter.tableName).note;
      
      events.note.index = admin.popup({
        title: '便签'
        ,shade: 0
        ,offset: [
          '41px'
          ,(mobile ? null : (othis.offset().left - 250) + 'px')
        ]
        ,anim: -1
        ,id: 'LAY_adminNote'
        ,skin: 'layadmin-note layui-anim layui-anim-upbit'
        ,content: '<textarea placeholder="内容"></textarea>'
        ,resize: false
        ,success: function(layero, index){
          var textarea = layero.find('textarea')
          ,value = note === undefined ? '便签中的内容会存储在本地，这样即便你关掉了浏览器，在下次打开时，依然会读取到上一次的记录。是个非常小巧实用的本地备忘录' : note;
          
          textarea.val(value).focus().on('keyup', function(){
            layui.data(setter.tableName, {
              key: 'note'
              ,value: this.value
            });
          });
        }
      })
    }
    
    //弹出关于面板
    ,about: function(){
      admin.popupRight({
        id: 'LAY_adminPopupAbout'
        ,success: function(){
          view(this.id).render('system/about')
        }
      });
    }
    
    //弹出更多面板
    ,more: function(){
      admin.popupRight({
        id: 'LAY_adminPopupMore'
        ,success: function(){
          view(this.id).render('system/more')
        }
      });
    }
    
    //返回上一页
    ,back: function(){
      history.back();
    }
    
    //主题设置
    ,setTheme: function(othis){
      var theme = setter.theme
      ,index = othis.data('index')
      ,nextIndex = othis.siblings('.layui-this').data('index');
      
      if(othis.hasClass(THIS)) return;
      
      othis.addClass(THIS).siblings('.layui-this').removeClass(THIS);
      
      if(theme.color[index]){
        theme.color[index].index = index
        admin.theme({
          color: theme.color[index]
        });
      }
    }
    
    //左右滚动页面标签
    ,rollPage: function(type, index){
      var tabsHeader = $('#LAY_app_tabsheader')
      ,liItem = tabsHeader.children('li')
      ,scrollWidth = tabsHeader.prop('scrollWidth')
      ,outerWidth = tabsHeader.outerWidth()
      ,tabsLeft = parseFloat(tabsHeader.css('left'));
      
      //右左往右
      if(type === 'left'){
        if(!tabsLeft && tabsLeft <=0) return;
        
        //当前的left减去可视宽度，用于与上一轮的页标比较
        var  prefLeft = -tabsLeft - outerWidth; 

        liItem.each(function(index, item){
          var li = $(item)
          ,left = li.position().left;
          
          if(left >= prefLeft){
            tabsHeader.css('left', -left);
            return false;
          }
        });
      } else if(type === 'auto'){ //自动滚动
        (function(){
          var thisLi = liItem.eq(index), thisLeft;
          
          if(!thisLi[0]) return;
          thisLeft = thisLi.position().left;
          
          //当目标标签在可视区域左侧时
          if(thisLeft < -tabsLeft){
            return tabsHeader.css('left', -thisLeft);
          }
          
          //当目标标签在可视区域右侧时
          if(thisLeft + thisLi.outerWidth() >= outerWidth - tabsLeft){
            var subLeft = thisLeft + thisLi.outerWidth() - (outerWidth - tabsLeft);
            liItem.each(function(i, item){
              var li = $(item)
              ,left = li.position().left;
              
              //从当前可视区域的最左第二个节点遍历，如果减去最左节点的差 > 目标在右侧不可见的宽度，则将该节点放置可视区域最左
              if(left + tabsLeft > 0){
                if(left - tabsLeft > subLeft){
                  tabsHeader.css('left', -left);
                  return false;
                }
              }
            });
          }
        }());
      } else {
        //默认向左滚动
        liItem.each(function(i, item){
          var li = $(item)
          ,left = li.position().left;

          if(left + li.outerWidth() >= outerWidth - tabsLeft){
            tabsHeader.css('left', -left);
            return false;
          }
        });
      }      
    }
    
    //向右滚动页面标签
    ,leftPage: function(){
      events.rollPage('left');
    }
    
    //向左滚动页面标签
    ,rightPage: function(){
      events.rollPage();
    }
    
    //关闭当前标签页
    ,closeThisTabs: function(){
      admin.closeThisTabs();
    }
    
    //关闭其它标签页
    ,closeOtherTabs: function(type){
      var TABS_REMOVE = 'LAY-system-pagetabs-remove';
      if(type === 'all'){
        $(TABS_HEADER+ ':gt(0)').remove();
        $(APP_BODY).find('.'+ TABS_BODY+ ':gt(0)').remove();
      } else {
        $(TABS_HEADER).each(function(index, item){
          if(index && index != admin.tabsPage.index){
            $(item).addClass(TABS_REMOVE);
            admin.tabsBody(index).addClass(TABS_REMOVE);
          }
        });
        $('.'+ TABS_REMOVE).remove();
      }
    }
    
    //关闭全部标签页
    ,closeAllTabs: function(){
      events.closeOtherTabs('all');
      location.hash = '';
    }
    
    //遮罩
    ,shade: function(){
      admin.sideFlexible();
    }
    ,nav: function(othis){
      var id = othis.data('id');
      layui.data('menuParam', {
        key:'id',
        value:id
      });
      events.closeOtherTabs('all');
      view('TPL_layout').refresh(function(){
        view().render('/'+othis.attr('lay-href')).then(function(res){
          var router = layui.router();
          var pathURL = admin.correctRouter(router.path.join('/'));
          var FILTER_TAB_TBAS = 'layadmin-layout-tabs';
          var matchTo
            ,tabs = $('#LAY_app_tabsheader>li');

          tabs.each(function(index){
            var li = $(this)
              ,layid = li.attr('lay-id');

            if(layid === pathURL){
              matchTo = true;
              admin.tabsPage.index = index;
            }
          });


          //如果未在选项卡中匹配到，则追加选项卡
          if(setter.pageTabs && pathURL !== '/'){
            if(!matchTo){
              $(APP_BODY).append('<div class="layadmin-tabsbody-item layui-show"></div>');
              admin.tabsPage.index = tabs.length;
              element.tabAdd(FILTER_TAB_TBAS, {
                title: '<span>'+ (res.title || '新标签页') +'</span>'
                ,id: pathURL
                ,attr: router.href
              });
            }
          }
          this.container = admin.tabsBody(admin.tabsPage.index);

          //定位当前tabs
          element.tabChange(FILTER_TAB_TBAS, pathURL);
          admin.tabsBodyChange(admin.tabsPage.index);
        });

      });
    }
  };
  
  //初始
  !function(){
    //主题初始化
    var local = layui.data(setter.tableName);
    local.theme && admin.theme(local.theme);
    
    //禁止水平滚动
    $body.addClass('layui-layout-body');
    
    //移动端强制不开启页面标签功能
    if(admin.screen() < 1){
      delete setter.pageTabs;
    }
    
    //不开启页面标签时
    if(!setter.pageTabs){
      container.addClass('layadmin-tabspage-none');
    }
    
    //低版本IE提示
    if(device.ie && device.ie < 10){
      view.error('IE'+ device.ie + '下访问可能不佳，推荐使用：Chrome / Firefox / Edge 等高级浏览器', {
        offset: 'auto'
        ,id: 'LAY_errorIE'
      });
    }
    
  }();
  
  admin.prevRouter = {}; //上一个路由
  admin.prevErrorRouter = {}; //上一个异常路由

    //左侧导航切换
    element.tab({
        headerElem: '#'+ SIDE_MENU1 +'>li'
        ,bodyElem: '#'+ SIDE_MENU2 +'>.layui-menu-item'
    });

    //监听侧边一级菜单切换
    element.on('tab(layadmin-system-menu)', function(obj){
        if(admin.screen() < 2){
            admin.sideFlexible('spread');
            admin.iconFlexible();
        }
    });
  
  //监听 hash 改变侧边状态
  admin.on('hash(side)', function(router){
    view('TPL_layout').refresh(function(){
      element.render('nav', 'layadmin-side-child'); //重新渲染子菜单
      admin.iconFlexible(); //根据二级菜单情况显示/隐藏icon
    });

    /*var path = router.path, getData = function(item){
      return {
        list: item.children('.layui-nav-child')
        ,name: item.data('name')
        ,jump: item.data('jump')
      }
    }
    ,sideMenu = $('#'+ SIDE_MENU)
    ,SIDE_NAV_ITEMD = 'layui-nav-itemed'
    
    //捕获对应菜单
    ,matchMenu = function(list){
      var pathURL = admin.correctRouter(path.join('/'));
      list.each(function(index1, item1){
        var othis1 = $(item1)
        ,data1 = getData(othis1)
        ,listChildren1 = data1.list.children('dd')
        ,matched1 = path[0] == data1.name || (index1 === 0 && !path[0]) 
        || (data1.jump && pathURL == admin.correctRouter(data1.jump));
        
        listChildren1.each(function(index2, item2){
          var othis2 = $(item2)
          ,data2 = getData(othis2)
          ,listChildren2 = data2.list.children('dd')
          ,matched2 = (path[0] == data1.name && path[1] == data2.name)
          || (data2.jump && pathURL == admin.correctRouter(data2.jump));
          
          listChildren2.each(function(index3, item3){
            var othis3 = $(item3)
            ,data3 = getData(othis3)
            ,matched3 = (path[0] ==  data1.name && path[1] ==  data2.name && path[2] == data3.name)
            || (data3.jump && pathURL == admin.correctRouter(data3.jump))
            
            if(matched3){
              var selected = data3.list[0] ? SIDE_NAV_ITEMD : THIS;
              othis3.addClass(selected).siblings().removeClass(selected); //标记选择器
              return false;
            }
            
          });

          if(matched2){
            var selected = data2.list[0] ? SIDE_NAV_ITEMD : THIS;
            othis2.addClass(selected).siblings().removeClass(selected); //标记选择器
            return false
          }
          
        });
        
        if(matched1){
          var selected = data1.list[0] ? SIDE_NAV_ITEMD : THIS;
          othis1.addClass(selected).siblings().removeClass(selected); //标记选择器
          return false;
        }
        
      });
    }
    
    //重置状态
    sideMenu.find('.'+ THIS).removeClass(THIS);
    
    //移动端点击菜单时自动收缩
    if(admin.screen() < 2) admin.sideFlexible();
    
    //开始捕获
    matchMenu(sideMenu.children('li'));*/

  });
  
  //监听侧边导航点击事件
  element.on('nav(layadmin-side-child)', function(elem){
    if(elem.siblings('.layui-nav-child')[0] && container.hasClass(SIDE_SHRINK)){
      admin.sideFlexible('spread');
      layer.close(elem.data('index'));
    };
  });
  
  //监听选项卡的更多操作
  element.on('nav(layadmin-pagetabs-nav)', function(elem){
    var dd = elem.parent();
    dd.removeClass(THIS);
    dd.parent().removeClass(SHOW);
  });
  
  //同步路由
  var setThisRouter = function(othis){
    var layid = othis.attr('lay-id')
    ,index = othis.index();
    // console.log(layid, index);
    
    admin.tabsBodyChange(index);
    location.hash = layid === setter.entry ? '/' : layid;
  }
  ,TABS_HEADER = '#LAY_app_tabsheader>li';
  
  //页面标签点击
  $body.on('click', TABS_HEADER, function(){
    var othis = $(this)
    ,index = othis.index();

    if(othis.hasClass(THIS)){
      delete admin.tabsPage.type;
    }else{
      admin.tabsPage.type = 'tab';
    }

    admin.tabsPage.index = index;
    
    //如果是iframe类型的标签页
    if(othis.attr('lay-attr') === 'iframe'){
      return admin.tabsBodyChange(index);
    };
    
    //单页标签页
    setThisRouter(othis);
    
    //执行resize事件，如果存在的话
    admin.runResize();
  });
  
  //监听 tabspage 删除
  element.on('tabDelete(layadmin-layout-tabs)', function(obj){
    var othis = $(TABS_HEADER+ '.layui-this');
    
    obj.index && admin.tabsBody(obj.index).remove();
    setThisRouter(othis);
    
    //移除resize事件
    admin.delResize();
  });
  
  //页面跳转
  $body.on('click', '*[lay-href]', function(){
    var othis = $(this)
    ,href = othis.attr('lay-href')
    ,router = layui.router();
    
    admin.tabsPage.elem = othis;
    //console.log(layui.admin.tabsPage);
    admin.prevRouter[router.path[0]] = router.href; //记录上一次各菜单的路由信息

    //执行跳转
    location.hash = admin.correctRouter(href);
  });
  
  //点击事件
  $body.on('click', '*[layadmin-event]', function(){
    var othis = $(this)
    ,attrEvent = othis.attr('layadmin-event');
    events[attrEvent] && events[attrEvent].call(this, othis);
  });
  
  //tips
  $body.on('mouseenter', '*[lay-tips]', function(){
    var othis = $(this);
    
    if(othis.parent().hasClass('layui-nav-item') && !container.hasClass(SIDE_SHRINK)) return;
    
    var tips = othis.attr('lay-tips')
    ,offset = othis.attr('lay-offset') 
    ,direction = othis.attr('lay-direction')
    ,index = layer.tips(tips, this, {
      tips: direction || 1
      ,time: -1
      ,success: function(layero, index){
        if(offset){
          layero.css('margin-left', offset + 'px');
        }
      }
    });
    othis.data('index', index);
  }).on('mouseleave', '*[lay-tips]', function(){
    layer.close($(this).data('index'));
  });
  
  //窗口resize事件
  var resizeSystem = layui.data.resizeSystem = function(){
    //layer.close(events.note.index);
    layer.closeAll('tips');
    
    if(!resizeSystem.lock){
      setTimeout(function(){
        admin.sideFlexible(admin.screen() < 2 ? '' : 'spread');
        delete resizeSystem.lock;
      }, 100);
    }
    
    resizeSystem.lock = true;
  }
  $win.on('resize', layui.data.resizeSystem);
  
  //接口输出
  exports('admin', admin);
});