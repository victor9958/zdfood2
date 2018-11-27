layui.define(['admin', 'form', 'table'], function (exports) {
    var $ = layui.$,
        admin = layui.admin,
        setter = layui.setter,
        form = layui.form,
        table = layui.table,
        element = layui.element,
        view = layui.view;

    var roleId, roleName;


    $(function () {
        getRoleList();

        //点击切换角色
        $(document).off('click', '.role-item');
        $(document).on('click', '.role-item', function () {
            $(this).addClass('active').siblings().removeClass('active');
            roleId = $(this).attr('data-id');
            roleName = $(this).attr('data-name');
            layui.data('roleParam', {key: 'id', value: roleId});
            layui.data('roleParam', {key: 'name', value: roleName});
            view('approve-tpl').refresh();
        });

        //取消设置
        $(document).off('click', '#module_cancel');
        $(document).on('click', '#module_cancel', function () {
            layui.index.render();
        });

        //获取角色列表
        function getRoleList() {
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/roles-and-approve',
                type: 'get',
                done: function (res) {
                    var html = '', active = '';
                    if (!roleId) {
                        roleId = res.data.roles[0].id;
                        roleName = res.data.roles[0].role_name;
                    }
                    if (res.data.roles.length > 0) {
                        $.each(res.data.roles, function (i, k) {
                            console.log(k.id)
                            console.log(roleId)
                            if (k.id == roleId) {
                                active = 'active';
                            } else {
                                active = '';
                            }
                            html += '<li class="list-item flex flex-align-center '+ active +'" data-id="'+ k.id +'"> ' +
                                '<span class="text flex-1">'+ k.role_name +'</span> ' +
                                '<svg class="icon goto" aria-hidden="true"> ' +
                                '<use xlink:href="#icon-fh"></use> ' +
                                '</svg></li>';
                        })
                        $('#roleList').html(html);
                        // roleId = $('.active').attr('data-id');
                        console.log(roleId)
                        // getUsers();
                        layui.data('roleParam', {key: 'id', value: roleId});
                        layui.data('roleParam', {key: 'name', value: roleName});
                    }

                    if(res.data.approve === 0){
                        $('.switchStatus').removeAttr('checked')
                    }else{
                        $('.switchStatus').attr('checked','')
                    }
                    form.render();
                }
            })
        }

        //全局开关
        form.on('switch(switchStatus)', function(data){
            var status = 0;
            if(data.elem.checked){
                status = 1;
            }else{
                status = 0;
            }
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/switch-status',
                type: 'put',
                data: {status:status},
                done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            })
        });

        // $(document).off('click','.layui-card');
        // $(document).on('click','.layui-card',function(){
        //
        // })


        // 提交设置
        form.on('submit(LAY-module-submit)',function(obj){
            var jobNumberInfo = {};
            var options = [];
            var structureInfo = {};
            var countCheck = 0,countRadio = 0;
            $(document).find('.module-card').each(function(){
                var elem = $(this).find('input[type="checkbox"]');
                var value = elem.prop('checked');
                var id = elem.attr('data-id');
                if(elem.attr('name') === 'jobNumberInfo' && value){
                    jobNumberInfo.status = 1
                }else if(elem.attr('name') === 'jobNumberInfo' && !value){
                    jobNumberInfo.status = 0
                }
                if(elem.attr('name') === 'structureInfo' && value){
                    structureInfo.status = 1
                }else if(elem.attr('name') === 'structureInfo' && !value){
                    structureInfo.status = 0
                }
                if(elem.attr('name').indexOf('options_status')==0 && value){
                    var id = elem.attr('data-id');
                    options[countCheck] = {id:id,status : 1};
                    countCheck ++;
                }else if(elem.attr('name').indexOf('options_status')==0 && !value){
                    var id = elem.attr('data-id');
                    options[countCheck] = {id:id,status : 0};
                    countCheck ++;
                }
            });
            $(document).find('.module-card').each(function(){
                var elem = $(this).find('input[type="radio"]');
                var value = elem.prop('checked');
                var id = elem.attr('data-id');
                if(elem.attr('name') === 'jobNumberInfo_radio' && value){
                    jobNumberInfo.require = 1
                }else if(elem.attr('name') === 'jobNumberInfo_radio' && !value){
                    jobNumberInfo.require = 0
                }
                if(elem.attr('name') === 'structureInfo_radio' && value){
                    structureInfo.require = 1
                }else if(elem.attr('name') === 'structureInfo_radio' && !value){
                    structureInfo.require = 0
                }
                if(elem.attr('name').indexOf('options_require')==0 && value){
                    options[countRadio]['require'] = 1;
                    countRadio ++;
                }else if(elem.attr('name').indexOf('options_require')==0 && !value){
                    options[countRadio]['require'] = 0;
                    countRadio ++;
                }
            })
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/role/'+ layui.data('roleParam').id +'/approve-tpl',
                type: 'put',
                data: {jobNumberInfo:jobNumberInfo,options:options,structureInfo:structureInfo},
                done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            })

        })
    })
    exports('setting/auditmodule', {})
});