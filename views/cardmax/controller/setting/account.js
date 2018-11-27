layui.define(function (exports) {
    layui.use(['form'], function () {
        var $ = layui.$,
            view = layui.view,
            form = layui.form,
            setter = layui.setter,
            admin = layui.admin;

        $(function () {
            $(document).off('click', '#resetmobile')
            $(document).on('click', '#resetmobile', function () {
                admin.popupCenter({
                    id: 'LAY_adminPopupResetMobile'
                    , title: '更换手机号'
                    , area: ['400px', '270px']
                    , success: function () {
                        view(this.id).render('setting/center/account/resetmobile');
                    }
                });
            })

            $(document).off('click', '#reset')
            $(document).on('click', '#reset', function () {
                admin.popupCenter({
                    id: 'LAY_adminPopupResetPassword'
                    , title: '修改密码'
                    , area: ['400px', '320px']
                    , success: function () {
                        view(this.id).render('setting/center/account/resetpassword');
                    }
                });
            })

            form.on('submit(LAY-reset-mobile)', function (obj) {
                //请求添加部门接口
                admin.req({
                    url: setter.baseUrl + '/system-set/account/mobile'
                    , type: 'put'
                    , data: obj.field
                    , done: function (res) {
                        view('old_mobile')
                        layer.msg(res.message, {icon: 1, time: 2000});
                        layer.close(admin.popup.index);//关闭右侧面板
                    }
                });
            });

            //关闭弹窗
            $(document).off('click', '#cancel');
            $(document).on('click', '#cancel', function () {
                layer.close(admin.popup.index);
            })

        })
    })
    exports('setting/account', {})
});