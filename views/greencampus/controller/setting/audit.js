layui.define(['admin', 'form', 'table', 'Sortable'], function (exports) {
    var $ = layui.$,
        admin = layui.admin,
        setter = layui.setter,
        form = layui.form,
        table = layui.table,
        element = layui.element,
        Sortable = layui.Sortable,
        view = layui.view;

    var roleId, departmentId, queryString, approve_tpl_id, downHref;
    var roleId_, departmentId_;
    var oldTplName;
    var sortList = [], idList = [];
    var statusList = [0];
    var reason, userList;
    var cols_examine = [
        {type: 'checkbox', fixed: 'left'}
        , {type: 'numbers', title: '序号', width: 80}
        , {field: 'name', title: '姓名', width: 100}
        , {field: 'job_number', title: '学号/工号', width: 140}
        , {field: 'mobile', title: '手机号', width: 120}
        , {field: 'role_name', title: '角色', width: 80}
        , {field: 'structureName', title: '组织架构'}
        , {title: '操作', width: 200, toolbar: '#option-examine'}
    ];
    var cols_examined = [
        {type: 'numbers', title: '序号', width: 80}
        , {field: 'name', title: '姓名', width: 100}
        , {field: 'job_number', title: '学号/工号', width: 140}
        , {field: 'mobile', title: '手机号', width: 120}
        , {field: 'role_name', title: '角色', width: 80}
        , {field: 'structureName', title: '组织架构'}
        , {field: 'status', title: '审核结果', templet: '#isPass'}
        , {title: '操作', toolbar: '#option-examined'}
    ];

    $(function () {
        form.render();
        getSelect('.role-option');
        getSelect('.role-list');
        tableData('examine', statusList, cols_examine);

        search('submit(LAY-search-examine)');
        search('submit(LAY-search-examined)');

        btnOption('tool(examineTable)');
        btnOption('tool(examinedTable)');

        element.on('tab(tab)', function (data) {
            form.render();
            var idx = data.index;
            approve_tpl_id = $('.content-right .layui-tab-title .layui-this').attr('lay-id');
            $('.role').html('');
            $('.department').html('');
            if (idx === 0) {
                statusList = [0];
                tableData('examine', statusList, cols_examine);
                getSelect('.role');
            } else if (idx === 1) {
                statusList = [1, 2];
                console.log(statusList)
                tableData('examined', statusList, cols_examined);
                getSelect('.role');
            } else {
                sortable(approve_tpl_id);
                sortList = [];
                $('#sortable-' + approve_tpl_id + ' li').each(function (i, k) {
                    sortList.push($(this).data('sort'));
                })
            }
        });

        //添加填写项
        $(document).off('click', '#addmodal');
        $(document).on('click', '#addmodal', function () {
            var count = $('.count-' + approve_tpl_id).html();
            var total = $('.total-' + approve_tpl_id).html();
            console.log(count)
            console.log(count)
            console.log(total)
            if (total - count > 0) {
                admin.popupCenter({
                    id: 'LAY_adminPopupAddModal'
                    , title: '添加填写项'
                    , area: ['520px', '500px']
                    , success: function () {
                        view(this.id).render('setting/user/audit/addmodal');
                    }
                });
            } else {
                layer.alert('最多只能添加3个哦', {icon: 5});
            }
        })


        //select身份监听事件
        form.on('select(role)', function (data) {
            var roleId = data.value;
            if (roleId != '') {
                getSelect('.department', '/user-manage/audit/department', statusList, roleId);
            } else {
                getSelect('.department', '/user-manage/audit/department', statusList);
            }
        });
        form.on('select(role_)', function (data) {
            var roleId = data.value;
            $('#roleId').val(roleId);
            form.render();
            if (roleId != '') {
                getSelect('#departmentId', '/user-manage/audit/department', [], roleId);
            }
        });

        //批量通过
        $(document).off('click', '#pass');
        $(document).on('click', '#pass', function () {
            if (!$(this).hasClass('layui-btn-disabled')) {
                userList = getUserId('examine');
                if (userList.length > 0) {
                    layer.confirm('是否确定通过审核？', {icon: 3, title: '提示'}, function (index) {
                        examine(1);
                    });
                } else {
                    layer.alert('请选择要操作的成员！', {icon: 5});
                }
            }
        })

        //批量不通过
        $(document).off('click', '#unpass');
        $(document).on('click', '#unpass', function () {
            if (!$(this).hasClass('layui-btn-disabled')) {
                userList = getUserId('examine');
                if (userList.length > 0) {
                    unPass();
                } else {
                    layer.alert('请选择要操作的成员！', {icon: 5});
                }
            }

        })

        //点击编辑模板名称
        $(document).off('click', '.edit-modal-name');
        $(document).on('click', '.edit-modal-name', function () {
            approve_tpl_id = $(this).data('id');
            oldTplName = $(this).siblings('.tplname').text();
            var content = '<div class="layui-form full-height centerLayer">' +
                '<div class="layui-form-item">' +
                '<label class="layui-form-label left">模板名称</label>' +
                '<div class="layui-input-block">' +
                '<input type="text" id="modal_name" autocomplete="off" value="' + oldTplName + '" placeholder="请输入模板名称" class="layui-input">' +
                '</div>' +
                '</div>' +
                '</div>';
            admin.popupCenter({
                id: 'LAY_adminModulename'
                , title: '修改模块名'
                , area: ['400px', '200px']
                , btn: ['确定', '取消']
                , content: content
                , yes: function () {
                    var tplName = $('#modal_name').val();
                    if (!tplName) {
                        $(this).val(oldTplName);
                        layer.msg('请输入模板名称！', {icon: 5, time: 1000});
                    } else {
                        $('.modal-name').hide();
                        resetModal(approve_tpl_id, tplName);
                    }
                }
            });
        })

        //添加模板
        var count;
        $(document).off('click', '#add-modal');
        $(document).on('click', '#add-modal', function () {
            count = $('.content-right .layui-tab-title li').length;
            if (count < 3) {
                admin.req({
                    url: setter.baseUrl + '/user-manage/audit/approve-tpl'
                    , type: 'post'
                    , data: {'tplName': '新增模板'}
                    , done: function (res) {
                        count++;
                        var tab_content = '<div class="m-b-10">' +
                            '<p class="layui-form-label" style="width: 500px;padding-left: 0;">' +
                            '<span class="ft-red">*</span>若成员在登录时出现操作或权限疑问，以方便联系管理员，请输入联系方式： </p>' +
                            '<div class="layui-input-inline"> <input type="text" id="contact" autocomplete="off" value="" placeholder="请输入联系方式" class="layui-input">' +
                            '</div></div>' +
                            '<div class="handle relative_">模板名称：<span class="tplname">新增模板</span>' +
                            '<svg class="icon edit-modal-name" aria-hidden="true"><use xlink:href="#icon-bianji"></use></svg>' +
                            '<div class="right"><span class="count-' + res.data + '">0</span>/<span class="total-' + res.data + '">3</span>' +
                            '<a href="javascript:;" id="addmodal" class="layui-btn layui-btn-sm" style="margin-left: 10px;">添加填写项</a></div></div>' +
                            '<div class="ui-state-default tb-mode"><div class="tb-cell width-40"></div><div class="tb-cell percent-25">名称</div>' +
                            '<div class="tb-cell percent-25">类型</div><div class="tb-cell">提示</div><div class="tb-cell width-80"></div></div>' +
                            '<ul class="sortable" id="sortable-' + res.data + '"></ul><div class="m-t-40 ft-center"><button class="layui-btn use-module" data-id="' + res.data + '">使用模板</button>' +
                            '<button class="layui-btn layui-btn-primary" id="delete-modal" data-id="' + res.data + '">删除模板</button></div>';
                        // $('.content-right .layui-tab-title').append('<li lay-id="'+ res.data +'">新增模板</li>');
                        // $('.content-right .layui-tab-content').append('<div class="layui-tab-item layui-tab-item-'+ res.data +'">'+ tab_content +'</div>');
                        element.tabAdd('tab-child', {
                            title: '新增模板'
                            , content: tab_content
                            , id: res.data
                        });
                        element.tabChange('tab-child', res.data);
                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                });
            } else {
                layer.alert('最多只能添加2个模板哦', {icon: 5});
            }
        })

        //添加选项
        form.on('submit(LAY-add-option-submit)', function (obj) {
            obj.field.approve_tpl_id = approve_tpl_id;
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/form-option'
                , type: 'post'
                , data: obj.field
                , done: function (res) {
                    layer.close(admin.popup.index);//关闭面板
                    getSingleModule();
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        });

        //删除模板
        $(document).off('click', '#delete-modal');
        $(document).on('click', '#delete-modal', function () {
            var id = $(this).data('id');
            layer.confirm('是否删除该模板？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + '/user-manage/audit/approve-tpl/' + id
                    , type: 'delete'
                    , done: function (res) {
                        element.tabDelete('tab-child', id);
                        var idx = $('.content-right .layui-tab-title .layui-this').index();
                        $('.layui-tab-item-' + id).remove();
                        $('.content-right .layui-tab-item').eq(idx - 1).addClass('layui-show').siblings().removeClass('layui-show');
                        approve_tpl_id = $('.content-right .layui-this').attr('lay-id');
                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                });
            });
        })

        //删除填写项
        $(document).off('click', '#delete-option');
        $(document).on('click', '#delete-option', function () {
            var id = $(this).data('id');
            layer.confirm('是否删除该填写项？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + '/user-manage/audit/form-option/' + id
                    , type: 'delete'
                    , done: function (res) {
                        getSingleModule();
                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                });
            });
        })

        //
        element.on('tab(tab-child)', function (data) {
            var idx = data.index;
            approve_tpl_id = $('.content-right .layui-tab-title li').eq(idx).attr('lay-id');
            sortable(approve_tpl_id);
            sortList = [];
            $('#sortable-' + approve_tpl_id + ' li').each(function (i, k) {
                sortList.push($(this).data('sort'));
            })
            console.log(sortList)
        });

        //导出
        form.on('submit(LAY-export)', function (obj) {
            roleId = obj.field.roleId;
            departmentId = obj.field.departmentId;
            queryString = obj.field.queryString;

            // console.log(setAuditDownHref())
            window.location.href = setAuditDownHref();
        });

        /*detail*/
        $(document).off('click', '#editchange');
        $(document).on('click', '#editchange', function () {
            $(this).hide();
            roleId_ = $('#role').data('id');
            departmentId_ = $('#department').data('id');
            console.log(roleId_);
            console.log(departmentId_);
            $('#LAY_adminPopupDetail .centerLayer').addClass('edit-form');
            getSelect('#roleId');
        })

        //保存并提交
        form.on('submit(LAY-save-pass)', function (obj) {
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/user/' + obj.field.id
                , type: 'put'
                , data: obj.field
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000});
                    layer.closeAll();
                    tableData('examine', statusList, cols_examine);
                }
            });
        });

        //使用模板
        $(document).off('click', '.use-module');
        $(document).on('click', '.use-module', function () {
            var id = $(this).data('id');
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/approve-tpl/' + id + '/use'
                , type: 'put'
                , done: function (res) {
                    view('modal-content').refresh(function () {
                        element.tabChange('tab-child', id);
                    });
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        })

        //取消使用模板
        $(document).off('click', '#cancel-tpl');
        $(document).on('click', '#cancel-tpl', function () {
            var id = $(this).data('id');
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/approve-tpl/' + id + '/cancel'
                , type: 'put'
                , done: function (res) {
                    view('modal-content').refresh(function () {
                        element.tabChange('tab-child', id);
                    });
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        })

        //保存号码
        $(document).on('change', '#contact', function () {
            var contact = $(this).val();
            approve_tpl_id = $(this).data('id');
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/tpl-contact/' + approve_tpl_id
                , type: 'put'
                , data: {'contact': contact}
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        });
        $(document).on('keypress', '#contact', function (event) {
            var tplName = $(this).val();
            var keynum = (event.keyCode ? event.keyCode : event.which);
            if (keynum == '13') {
                $(this).change();
            }
        });

        //关闭弹窗
        $(document).off('click', '#cancel');
        $(document).on('click', '#cancel', function () {
            layer.close(admin.popup.index);
        })

        //监听按钮操作
        function btnOption(even) {
            table.on(even, function (obj) {
                var data = obj.data;
                layui.data('temporary', {
                    key: 'temporary'
                    , value: {'id': data.id, 'statusList': statusList}
                });
                if (obj.event === 'detail') {
                    // admin.popupCenter({
                    //     id: 'LAY_adminPopupDetail'
                    //     , title: '详情'
                    //     , area: ['400px', '500px']
                    //     , success: function () {
                    //         view(this.id).render('setting/user/audit/detail');
                    //     }
                    // });
                    location.hash = '/setting/user/audit/detail'
                } else if (obj.event === 'pass') {
                    userList = [data.id];
                    layer.confirm('是否确定通过审核？', {icon: 3, title: '提示'}, function (index) {
                        examine(1);
                    });
                } else {
                    userList = [data.id];
                    unPass();
                }
            });

        }

        //提交搜索
        function search(even) {
            form.on(even, function (obj) {
                roleId = obj.field.roleId;
                departmentId = obj.field.departmentId;
                queryString = obj.field.queryString;
                if (even == 'submit(LAY-search-examine)') {
                    tableData('examine', statusList, cols_examine);
                } else {
                    if (obj.field.statusList == '') {
                        statusList = [1, 2];
                    } else {
                        statusList = [obj.field.statusList];
                    }
                    tableData('examined', statusList, cols_examined);
                }
            });
        }

        //获取表格中选中成员id
        function getUserId(elem) { //elem:table的id，例如'examine'
            var userId = [];
            var checkStatus = table.checkStatus(elem), data = checkStatus.data;
            for (var i in data) {
                userId.push(data[i].id);
            }
            return userId;
        }

        //表格数据
        function tableData(tableName, statusList, cols_) {//status_:0待审核 1通过 2不通过
            table.render({
                id: tableName
                , elem: '#' + tableName
                , headers: {'Authorization': layui.data(setter.tableName)[setter.request.tokenName]}
                , url: setter.baseUrl + '/user-manage/audit/list' //模拟接口
                , page: setter.pageTable
                , cellMinWidth: 100
                , loading: false
                , limit: setter.limit
                , response: setter.responseTable
                , cols: [cols_]
                , where: {
                    statusList: statusList
                    , roleId: roleId
                    , queryString: queryString
                }
                , skin: 'line'
            });
        }

        //获取表单身份初始数据
        function getSelect(elem) {
            var option = '';
            admin.req({
                url: setter.baseUrl + '/user-manage/audit/roles'
                , type: 'get'
                , data: {}
                , done: function (res) {
                    option = '<option value="">请选择身份</option>';
                    $.each(res.data, function (i, k) {
                        if (k.id == roleId_) {
                            option += '<option value="' + k.id + '" selected>' + k.role_name + '</option>';
                        } else {
                            option += '<option value="' + k.id + '">' + k.role_name + '</option>';
                        }
                    })
                    $(elem).html(option);
                    if (!$('#roleId').val()) {
                        $('#roleId').val(layui.data('roleParam').id);
                    }

                    form.render();
                }
            });
        }

        //审核
        function examine(passOrReject) {//passOrReject: 1通过  2不通过
            admin.req({
                url: setter.baseUrl + '/user-manage/audit'
                , type: 'put'
                , data: {'userList': userList, 'passOrReject': passOrReject, 'reason': reason}
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000});
                    layer.close(admin.popup.index);
                    tableData('examine', statusList, cols_examine);
                }
            });
        }

        //生成导出地址
        function setAuditDownHref() {
            downHref = setter.baseUrl + '/user-manage/audit/export?';
            $.each(statusList, function (i, k) {
                downHref += 'statusList[]=' + k + '&';
            })
            downHref += 'roleId=' + roleId + '&departmentId=' + departmentId + '&queryString=' + queryString;
            return downHref;
        }

        function unPass() {
            admin.popupCenter({
                id: 'LAY_adminExamine'
                , title: '审核操作'
                , area: ['400px', '260px']
                , success: function () {
                    view(this.id).render('setting/user/audit/examine');
                }
            });
        }

        form.on('submit(LAY-examine-submit)', function (obj) {
            reason = obj.field.reason;
            examine(2);
        });

    })
    exports('setting/audit', {})
});