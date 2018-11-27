layui.define(['table', 'form'], function (exports) {
    var $ = layui.$,
        setter = layui.setter,
        admin = layui.admin,
        view = layui.view,
        form = layui.form,
        element = layui.element,
        table = layui.table;

    $(function () {
        var roleId = '', roleName;
        getRoleList();

        //点击切换角色
        $(document).off('click', '.role-item');
        $(document).on('click', '.role-item', function () {
            $(this).addClass('active').siblings().removeClass('active');
            roleId = $(this).attr('data-id');
            roleName = $(this).attr('data-name');
            layui.data('roleParam', {key: 'id', value: roleId});
            layui.data('roleParam', {key: 'name', value: roleName});
            getModule();
            //getRoleModule(id);
        });

        //关闭弹窗
        $(document).off('click', '#cancel');
        $(document).on('click', '#cancel', function () {
            layer.close(admin.popup.index);
        });

        //点击角色置顶
        $(document).off('click', '#icon-zhiding');
        $(document).on('click', '#icon-zhiding', function (e) {
            e.stopPropagation();
            var current = $(this).parents('.role-item');
            var id = current.attr('data-id');
            // var idx = current.index();
            layer.confirm('是否对该角色进行置顶操作？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + '/user-manage/role-top/' + id,
                    type: 'put',
                    done: function (res) {
                        getRoleList();
                        layer.msg(res.message, {
                            icon: 1,
                            time: 1000
                        });
                    }
                });
            });
        });

        //搜索用户
        form.on('submit(LAY-search-users)', function (obj) {
            getUsers(obj.field.queryString)
        });

        //点击在我的模块中删除模块
        $(document).off('click', '#myModule .operation');
        $(document).on('click', '#myModule .operation', function () {
            var id = $(this).parent().attr('data-id');
            $(this).parent().remove();
            var target = $('#allModule .module-item').eq(indexInAllModule(id)).find('.operation').toggleClass('plus').toggleClass('reduce').text('+');
        });

        //点击在全部模块中删减模块
        $(document).off('click', '#allModule .operation');
        $(document).on('click', '#allModule .operation', function () {
            var id = $(this).parent().attr('data-id');
            if ($(this).hasClass('plus')) {
                $(this).toggleClass('plus').toggleClass('reduce').text('-');
                $(this).parent().clone().appendTo('#myModule');
            } else {
                $(this).toggleClass('reduce').toggleClass('plus').text('+');
                $('#myModule .module-item').eq(indexInMyModule(id)).remove();
            }
        });

        //点击编辑模板
        $(document).off('click', '#edit_module');
        $(document).on('click', '#edit_module', function () {
            $(this).parents('.layui-tab-item').addClass('edit-content');
        })

        //点击确认修改模块
        $(document).off('click', '#add_module');
        $(document).on('click', '#add_module', function () {
            var moduleChangeArr = [];
            $('#myModule .module-item').each(function (index, item) {
                moduleChangeArr.push($(item).attr('data-id'));
            });
            console.log(moduleChangeArr);
            admin.req({
                url: setter.baseUrl + '/user-manage/role/' + roleId + '/sync-module',
                type: 'put',
                data: {'moduleList': moduleChangeArr},
                done: function (res) {
                    getModule();
                    $('.edit-content').removeClass('edit-content');
                    layer.msg(res.message, {
                        icon: 1,
                        time: 2000
                    });
                }
            });
        });

        //点击取消修改模块
        $(document).off('click', '#cancel_add');
        $(document).on('click', '#cancel_add', function () {
            getModule();
            $('.edit-content').removeClass('edit-content');
        });

        //批量删除成员
        $(document).off('click', '#deleteUser');
        $(document).on('click', '#deleteUser', function () {
            var idList = getMemberId();
            if (idList.length > 0) {
                deleteUser(idList)
            } else {
                layer.alert('请选择要操作的成员！', {icon: 5});
            }
        })

        //添加成员
        // $(document).off('click', '#addUser');
        // $(document).on('click', '#addUser', function () {
        //     admin.popupCenter({
        //         id: 'LAY_adminPopupAddUser'
        //         , title: '添加用户'
        //         , area: ['800px', '520px']
        //         , success: function () {
        //             view(this.id).render('setting/user/role/addUser');
        //         }
        //     });
        // })

        //提交添加成员
        // form.on('submit(LAY-add-user-submit)', function (obj) {
        //     obj.field.roleId = roleId;
        //     obj.field.extends = {};
        //     $('*[name="extends"]').each(function (i, k) {
        //         obj.field.extends[k.id] = k.value;
        //     })
        //     $('*[name="structureId"]').each(function (i, k) {
        //         if (k.value != '') {
        //             obj.field.structureId = k.value;
        //         }
        //     })
        //     console.log(obj.field)
        //     admin.req({
        //         url: setter.baseUrl + '/user-manage/role/user'
        //         , type: 'post'
        //         , data: obj.field
        //         , done: function (res) {
        //             //关闭右侧面板
        //             layer.close(admin.popup.index);
        //             getUsers();
        //             layer.msg(res.message, {icon: 1, time: 2000});
        //         }
        //     });
        // });

        //批量删除成员
        $(document).off('click', '#setDept');
        $(document).on('click', '#setDept', function () {
            var idList = getMemberId();
            if (idList.length > 0) {
                admin.popupCenter({
                    id: 'LAY_adminPopupSetDept'
                    , title: '设置部门'
                    , area: ['400px', '420px']
                    , success: function () {
                        view(this.id).render('setting/user/role/setDept');
                    }
                });
            } else {
                layer.alert('请选择要操作的成员！', {icon: 5});
            }
        })

        //提交调整部门
        form.on('submit(LAY-set-dept-submit)', function (obj) {
            var userList = getMemberId();
            obj.field.userList = userList;
            $('*[name="structureId"]').each(function (i, k) {
                if (k.value != '') {
                    obj.field.structureId = k.value;
                }
            })
            admin.req({
                url: setter.baseUrl + '/user-manage/role/user/structure'
                , type: 'put'
                , data: obj.field
                , done: function (res) {
                    //关闭右侧面板
                    layer.close(admin.popup.index);
                    getUsers();
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        });

        //导入成员
        $(document).off('click', '#importUser');
        $(document).on('click', '#importUser', function () {
            admin.popupCenter({
                id: 'LAY_adminPopupImportUser'
                , title: '导入用户'
                , area: ['600px', '370px']
                , success: function () {
                    view(this.id).render('setting/user/role/importUser');
                }
            });
        })

        //
        element.on('tab(tabDetail)', function (data) {
            var idx = data.index;
            if (idx === 1) {
                getModule();
            }
        });

        //我的模块中查找索引
        function indexInMyModule(id) {
            var moduleArr = [];
            $('#myModule .module-item').each(function (index, item) {
                moduleArr.push($(item).attr('data-id'));
            });
            return moduleArr.indexOf(id)
        }

        //全部模块中查找索引
        function indexInAllModule(id) {
            var moduleArr = [];
            $('#allModule .module-item').each(function (index, item) {
                moduleArr.push($(item).attr('data-id'));
            });
            return moduleArr.indexOf(id)
        }

        //getRoleModule(1);
        function changeRoleName(id, name) {
            admin.req({
                type: 'put',
                url: setter.baseUrl + '/user-manage/role/' + id,
                data: {
                    roleName: name
                },
                done: function (res) {
                    // $('.changeName').blur();
                    // $('.rolename').html(name);
                    getRoleList();
                    layer.msg(res.message, {
                        icon: 1,
                        time: 2000
                    });
                    layer.close(admin.popup.index);//关闭面板
                }
            })
        }

        function getRoleModule(id) {
            admin.req({
                url: setter.baseUrl + '/user-manage/role/' + id + '/module-department',
                data: {
                    id: id
                },
                done: function (res) {
                    console.log(res.data);
                    appendModule(res.data.module);
                }
            });
        }

        function appendModule(data) {
            var html = '';
            $.each(data, function (index, item) {
                html += '<li class="item">' + item.module_name + '</li>'
            });
            $('.module-list').html(html);
        }

        //模块列表
        function getModule() {
            admin.req({
                url: setter.baseUrl + '/user-manage/role/' + roleId + '/module',
                type: 'get',
                done: function (res) {
                    var mymodule = '', allmodule = '';
                    if (res.data.module.myModule.length > 0) {
                        $.each(res.data.module.myModule, function (i, k) {
                            mymodule += '<div class="module-item ft-center" data-id="' + k.id + '"> ' +
                                '<p class="icon-block ' + k.icon_color + '">' +
                                '<svg class="icon" aria-hidden="true"><use xlink:href="#' + k.icon + '"></use></svg></p>' +
                                '<span class="text">' + k.module_name + '</span> ' +
                                '<span class="operation reduce">-</span> ' +
                                '</div>';
                        })
                    } else {
                        $('.layui-tab-item').addClass('edit-content');
                    }

                    $('#myModule').html(mymodule);

                    $.each(res.data.module.allModule, function (i, k) {
                        var operation = '', moduleitem = '';
                        $.each(k.modules, function (i2, k2) {
                            if (k2.isExist) {
                                operation = '<span class="operation reduce">-</span>';
                            } else {
                                operation = '<span class="operation plus">+</span>';
                            }
                            moduleitem += '<div class="module-item ft-center" data-id="' + k2.id + '"> ' +
                                '<p class="icon-block ' + k2.icon_color + '">' +
                                '<svg class="icon" aria-hidden="true"><use xlink:href="#' + k2.icon + '"></use></svg></p> ' +
                                '<span class="text">' + k2.module_name + '</span>' + operation +
                                '</div>';
                        })
                        allmodule += '<div class="module-title pull-left">' + k.title + '</div><div class="pull-left">' + moduleitem + '</div>';
                    })
                    $('#allModule').html(allmodule);
                }
            });
        }

        //获取角色列表
        function getRoleList() {
            admin.req({
                url: setter.baseUrl + '/user-manage/roles',
                type: 'get',
                done: function (res) {
                    var html = '', active = '';
                    if (!roleId) {
                        roleId = res.data[0].id;
                        roleName = res.data[0].role_name;
                    }
                    if (res.data.length > 0) {
                        $.each(res.data, function (i, k) {
                            if (k.id == roleId) {
                                active = 'active';
                            } else {
                                active = '';
                            }
                            html += '<li class="role-item ' + active + '" data-id="' + k.id + '" data-name="' + k.role_name + '">' +
                                '<span class="text">' + k.role_name + '</span>' +
                                '<div class="icon-btn-box pull-right">' +
                                '<svg class="icon" id="icon-zhiding" aria-hidden="true">' +
                                '<use xlink:href="#icon-zhiding"></use>' +
                                '</svg>' +
                                '</div>' +
                                '</li>';
                        })
                        $('.role-list').html(html);
                        // roleId = $('.active').attr('data-id');
                        console.log(roleId)
                        getUsers();
                        layui.data('roleParam', {key: 'id', value: roleId});
                        layui.data('roleParam', {key: 'name', value: roleName});
                    }

                }
            })
        }

        //获取角色用户列表
        function getUsers(queryString) {
            layer.closeAll('tips');
            table.render({
                id: 'users'
                , elem: '#users'
                , url: setter.baseUrl + '/user-manage/role/' + roleId + '/users' //模拟接口
                , page: setter.pageTable
                , cellMinWidth: 80
                , loading: false
                , limit: setter.limit
                , response: setter.responseTable
                , cols: [[
                    {type: 'checkbox', fixed: 'left'}
                    , {type: 'numbers', title: '序号'}
                    , {field: 'name', title: '姓名', width: 90, templet: '#is_admin'}
                    , {field: 'mobile', title: '手机号', width: 120}
                    , {field: 'job_number', title: '学号/工号', minWidth: 100}
                    , {field: 'roleName', title: '角色', minWidth: 80}
                    , {field: 'structureName', title: '组织架构', minWidth: 120}
                    , {title: '操作', width: 150, toolbar: '#userOption'}
                ]]
                , where: {
                    'queryString': queryString
                }
                , skin: 'line'
            });
        }

        // 表格按钮操作
        table.on('tool(usersTable)', function (obj) {
            var data = obj.data;
            if (obj.event === 'detail') {
                layui.data('currentId', {
                    key: 'id'
                    , value: data.id
                });
                // admin.popupCenter({
                //     id: 'LAY_adminPopupDetail'
                //     , title: '用户详情'
                //     , area: ['800px', '520px']
                //     , success: function () {
                //         view(this.id).render('setting/user/role/detail');
                //     }
                // });
                location.hash = '/setting/user/role/detail'
            } else if (obj.event === 'delete') {
                deleteUser([data.id]);
            }
        });

        //修改用户信息
        // form.on('submit(LAY-edit-user-submit)', function (obj) {
        //     obj.field.extends = {};
        //     $('*[name="extends"]').each(function (i, k) {
        //         obj.field.extends[k.id] = k.value;
        //     })
        //     $('*[name="structureId"]').each(function (i, k) {
        //         if (k.value != '') {
        //             obj.field.structureId = k.value;
        //         }
        //     })
        //     console.log(obj.field)
        //     admin.req({
        //         url: setter.baseUrl + '/user-manage/role/user/' + layui.data('currentId').id
        //         , type: 'put'
        //         , data: obj.field
        //         , done: function (res) {
        //             //关闭右侧面板
        //             layer.close(admin.popup.index);
        //             getUsers();
        //             layer.msg(res.message, {icon: 1, time: 2000});
        //         }
        //     });
        // });

        //获取表格中选中成员id
        function getMemberId() {
            var memberId = [];
            var checkStatus = table.checkStatus('users'), data = checkStatus.data;
            for (var i in data) {
                memberId.push(data[i].id);
            }
            return memberId;
        }

        // 删除用户
        function deleteUser(idList) {
            layer.confirm('是否删除选中成员？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + '/user-manage/role/batch-users'
                    , type: 'delete'
                    , data: {'idList': idList}
                    , done: function (res) {
                        getUsers();
                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                });
            });
        }

        //确认上传文件
        $(document).off('click', '#importConfirm');
        $(document).on('click', '#importConfirm', function () {
            var excelPath = $('#userExcelPath').val();
            if (!excelPath) {
                layer.alert('请先上传需要导入的文件！', {icon: 5});
            } else {
                admin.req({
                    url: setter.baseUrl + '/user-manage/role/user-import' //实际使用请改成服务端真实接口
                    , type: 'post'
                    , data: {'excelPath': excelPath}
                    , done: function (res) {
                        count = 0;
                        layer.close(admin.popup.index);
                        layer.msg(res.message, {icon:1, time: 2000});
                        getUsers()
                    }
                });
            }
        })

        var selectVal = '';
        form.on('select(structureSet)', function (data) {
            console.log($(data.elem).text().trim()); //得到select原始DOM对象
            console.log(data.value); //得到被选中的值
            console.log(data.othis); //得到美化后的DOM对象

            var currentVal = $(data.elem).text().trim();
            if(currentVal != selectVal || !data.value){
                $(data.elem).parents('.layui-form-item').nextAll().find('select').html('');
                if(data.value != ''){
                    admin.req({
                        url: setter.baseUrl + '/user-manage/roles/structure/'+ data.value +'/sons'
                        , type: 'get'
                        , data: {}
                        , done: function (res) {
                            if(res.data.length > 0){
                                var html = '<option value="">请选择</option>';
                                $.each(res.data,function(index,item){
                                    html += '<option value="'+ item.id +'">'+ item.name +'</option>';
                                })
                                console.log(html)
                                $(data.elem).parents('.layui-form-item').next().find('select').html(html);
                                form.render();
                            }
                            // layer.msg(res.message, {icon: 1, time: 1000});
                        }
                    });
                }
                form.render();
            }
        });

        $(document).off('click','.layui-form-select')
        $(document).on('click','.layui-form-select',function(){
            selectVal = $(this).find('input').val();
            console.log(selectVal)
        })

        //导出
        form.on('submit(LAY-export)', function (obj) {
            queryString = obj.field.queryString;
            // window.location.href = setAuditDownHref();
            window.open(setAuditDownHref())
        });

        //生成导出地址
        function setAuditDownHref() {
            downHref = setter.baseUrl + '/user-manage/role/user/export?';
            downHref += 'roleId=' + roleId + '&queryString=' + queryString;
            return downHref;
        }

    });
    exports('setting/role', {})
})