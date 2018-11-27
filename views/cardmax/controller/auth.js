layui.define('form', function (exports) {
    var $ = layui.$
        , setter = layui.setter
        , view = layui.view
        , admin = layui.admin
        , form = layui.form
        , router = layui.router()
        , search = router.search;


    $(function () {
        form.render();
        layer.closeAll();

        $('.layui-form').keypress(function (event) {
            var keynum = (event.keyCode ? event.keyCode : event.which);
            if (keynum == '13') {
                $('.layui-show .login-btn').off();
                $('.layui-show .login-btn').click();
                $('#forget').off();
                $('#forget').click();
            }
        });
        login('submit(LAY-password-login)','/auth/logon','post');
        login('submit(LAY-mobile-login)','/auth/logon','post');
        login('submit(LAY-setpassword-login)','/auth/password','put');

        //登录提交
        function login(even,url,type) {
            form.on(even, function (obj) {
                //请求登入接口
                admin.req({
                    url: setter.baseUrl + url
                    , data: obj.field
                    , type: type
                    , done: function (res) {
                        //请求成功后，写入 access_token
                        layui.data(setter.tableName, {
                            key: setter.request.tokenName
                            , value: 'Bearer '+res.data.access_token
                        });
                        layui.data('allMenu', {
                            key: 'menu'
                            , value: res.data.allMenu
                        });
                        layui.data('userMenu', {
                            key: 'menu'
                            , value: res.data.userMenu
                        });
                        layui.data('userNode', {
                            key: 'menu'
                            , value: res.data.userNode
                        });
                        layui.data('password', {
                            key: 'hasPassword'
                            , value: res.data.hasPassword
                        });

                        admin.req({
                            url: setter.baseUrl + '/auth/base-menu'
                            , done: function (res) {
                                layui.data('menuParam', {
                                    key: 'id',
                                    value: res.data[0].id
                                });
                            }
                        })

                        //登入成功的提示与跳转
                        layer.msg('登录成功', {
                            icon: 1
                            , time: 1000
                        }, function () {
                            location.hash = search.redirect ? decodeURIComponent(search.redirect) : '/';
                        });
                    }
                });

            });
        }

        form.verify({
            password: function(value,item){
                password = value;
                if(!/^[\S]{6,12}$/.test(value)){
                    return '密码必须6到12位，且不能出现空格';
                }
            }
            ,repassword: function(value, item){ //value：表单的值、item：表单的DOM对象
                if(value != password){
                    return '确认密码与新密码不符';
                }
            }
        });

        // 点击获取验证码
        var verifyTag = true;
        var phone_test = /^1[2-9]+\d{9}$/;
        $(document).off('click', '.get-verify');
        $(document).on('click', '.get-verify', function (e) {
            e.preventDefault();
            if (verifyTag) getCode(this);
            else return;
        });

        //倒计时
        function timeCountDown(elem, seconds) {
            verifyTag = false;
            if (typeof seconds == undefined || seconds == null) seconds = 60;
            $(elem).text('已发送 ' + seconds + 's').addClass('disabled');
            timeCount = setTimeout(function () {
                if (seconds > 0) {
                    seconds--;
                    $(elem).text('已发送 ' + seconds + 's').addClass('disabled');
                    timeCountDown(elem, seconds);
                } else {
                    $(elem).removeClass('disabled').text('获取验证码');
                    verifyTag = true;
                    clearTimeout(timeCount)
                }
            }, 1000);
        }

        function getCode(elem) {
            mobile = $('#mobile').val();
            if (!phone_test.test(mobile)) {
                layer.msg('请输入正确的手机号', {
                    icon: 5
                });
            } else {
                timeCountDown(elem);
                admin.req({
                    url: setter.baseUrl + '/auth/sms-code' //实际使用请改成服务端真实接口
                    // , data: {'mobile': mobile}
                    , type: 'get'
                    , done: function (res) {
                        layer.msg('已成功获取验证码', {
                            icon: 1
                            , time: 1000
                        });
                    }
                    ,error:function(){
                        $(elem).removeClass('disabled').text('获取验证码');
                    }
                });
            }
        }
    })


    //对外暴露的接口
    exports('auth', {});
});