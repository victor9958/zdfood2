layui.define(['admin', 'upload', 'form', 'laydate'], function (exports) {
    var $ = layui.$,
        admin = layui.admin,
        setter = layui.setter,
        view = layui.view,
        upload = layui.upload,
        laydate = layui.laydate,
        element = layui.element,
        form = layui.form;

    $(function () {
        var current_id = '';
        var current = new Date();
        current = current.getFullYear() + '-' + (current.getMonth() + 1) + '-' + current.getDate() + ' ' + current.getHours() + ':' + current.getMinutes() + ':' + current.getSeconds();

        element.on('tab(applicationtTab)', function (data) {
            var idx = data.index;
            if (idx === 1) {
                //公告遮罩
                current_id = $('.child-tab-content>.layui-show').find('.layui-btn').data('id');
                layDate('#date_' + current_id);
                save('submit(LAY-notice-banner' + current_id + ')');
                var cover_notice = document.getElementById("cover_notice");
                var context_notice = cover_notice.getContext("2d");
                context_notice.clearRect(0, 0, cover_notice.width, cover_notice.height);//抠出一个矩形区域
                context_notice.fillStyle = "rgba(0,0,0,0.3)";//设置填充色（可以是渐变色或半透明色）
                context_notice.fillRect(0, 0, cover_notice.width, cover_notice.height);//填充背景我色
                context_notice.clearRect(12, 149, 351, 28);//抠出一个矩形区域
            } else if (idx === 2) {
                //banner遮罩
                uploadImg();
                current_id = $('.child-tab-content>.layui-show').find('.layui-btn').data('id');
                layDate('#date_' + current_id);
                save('submit(LAY-notice-banner' + current_id + ')');
                var cover_banner = document.getElementById("cover_banner");
                var context_banner = cover_banner.getContext("2d");
                context_banner.clearRect(0, 0, cover_banner.width, cover_banner.height);//抠出一个矩形区域
                context_banner.fillStyle = "rgba(0,0,0,0.3)";//设置填充色（可以是渐变色或半透明色）
                context_banner.fillRect(0, 0, cover_banner.width, cover_banner.height);//填充背景我色
                context_banner.clearRect(12, 186, 351, 92);//抠出一个矩形区域
            }
        });

        element.on('tab(bannerTab)', function (data) {
            var idx = data.index;
            current_id = $(data.elem[0]).find('.layui-show .layui-btn').data('id');
            layDate('#date_' + current_id);
            save('submit(LAY-notice-banner' + current_id + ')');
        });

        //编辑模块名称
        var module_id;
        $(document).off('click', '.editmodulename');
        $(document).on('click', '.editmodulename', function () {
            var name = $(this).data('name');
            module_id = $(this).data('id');
            var content = '<div class="layui-form full-height centerLayer">' +
                '<div class="layui-form-item">' +
                '<label class="layui-form-label left">模块名称</label>' +
                '<div class="layui-input-block">' +
                '<input type="text" id="moduleName" autocomplete="off" value="' + name + '" placeholder="请输入模块名称" class="layui-input">' +
                '</div>' +
                '</div>' +
                '</div>';
            admin.popupCenter({
                id: 'LAY_adminModuleName'
                , title: '修改模块名称'
                , area: ['400px', '200px']
                , btn: ['确定', '取消']
                , content: content
                , yes: function () {
                    var name = $('#moduleName').val();
                    console.log(name)
                    if (!name) {
                        layer.msg('请输入模块名称', {icon: 5, time: 1000});
                    } else {
                        reset(name);
                    }
                }
            });
        })

        radio_check('radio(notice_time)');

        //radio事件
        var time_type = 1, notice_date, notice_time;

        function radio_check(filter) {
            form.on(filter, function (data) {
                if (data.value == ('timeout')) {
                    time_type = 2;
                    $('#timeout_' + current_id).siblings('.layui-inline').removeClass('hide');
                } else {
                    time_type = 1;
                    $('#timeout_' + current_id).siblings('.layui-inline').addClass('hide');
                    $('#date_' + current_id).val('');
                }
            });
        }

        //设置日期
        function layDate(elem_) {
            laydate.render({
                elem: elem_
                , type: 'datetime'
                // ,min: current
                , btns: ['clear', 'confirm']
            });
        }

        //提交编辑模块名称
        function reset(name) {
            admin.req({
                url: setter.baseUrl + '/system-set/module/' + module_id
                , type: 'put'
                , data: {'moduleAliasName': name}
                , done: function (res) {
                    view('app_content').refresh();
                    layer.msg(res.message, {icon: 1, time: 2000});
                    layer.close(admin.popup.index);//关闭面板
                }
            });
        }

        //保存设置
        function save(filter) {
            form.on(filter, function (obj) {
                delete obj.field.file;
                if (obj.field.status === 'on') {
                    obj.field.status = 1;
                } else {
                    obj.field.status = 0;
                }
                if (!obj.field.changeTime && time_type == 2) {
                    layer.alert('当前为定时更换，请选择时间', {icon: 5});
                } else if (new Date(obj.field.changeTime) < new Date(current)) {
                    layer.alert('定时更换时间不能小于当前时间', {icon: 5});
                } else {
                    admin.req({
                        url: setter.baseUrl + '/system-set/placard/' + $(this).data("id")
                        , type: 'put'
                        , data: obj.field
                        , done: function (res) {
                            layer.close(admin.popup.index);//关闭面板
                            // view('app_content').refresh();
                            layer.msg(res.message, {icon: 1, time: 2000});
                        }
                    });
                }
            });
        }

        function uploadImg() {
            upload.render({
                elem: '.upload_img'
                // ,field: 'content'
                , size: 2048
                , url: setter.baseUrl + '/public/oss-upload'
                , before: function (obj) {
                    //预读本地文件示例，不支持ie8
                    current_id = $(this)[0].item.data('id');
                }
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000})
                    $('.banner_img_' + current_id).attr('src', res.data.ossUrl); //图片链接（base64）
                    $('.content_' + current_id).val(res.data.ossUrl);
                }
            });
        }
    })
    exports('setting/application', {})
});