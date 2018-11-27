layui.define(['admin', 'table', 'Sortable', 'zTree'], function (exports) {
    var $ = layui.$,
        admin = layui.admin,
        table = layui.table,
        setter = layui.setter,
        view = layui.view,
        Sortable = layui.Sortable,
        form = layui.form;

    var deptId, deptName, baseDeptId, treeSelect, downHref;
    var dept = layui.data('currentDept').dept,
        zNodes = [],//组织架构数据
        zTreeObj = [],//组织架构树
        zTreeObj_ = [],//调整部门的树
        zTreeObj_check = [],//check树
        idList = [],//部门编号顺序
        sortList = [],//排序数组
        nameList = [],//部门名称数组
        sortableObj;//
    var deptIdList = [],//已选中的需要操作的部门数组
        departmentId;//单个部门
    var flag = true;
    $(function () {
        // var obj = {rule: "setting/user/organization/add_dept", method: "GET"};
        // console.log(layui.data('userNode').menu[0]);
        // console.log(obj);
        // console.log(JSON.stringify(layui.data('userNode').menu).indexOf(JSON.stringify(obj)));
        getTree('result', 'initPage');

        //拖拽排序
        $(document).off('click', '.sort');
        $(document).on('click', '.sort', function () {
            $('#sortable').removeClass('sortable');
            $('.option-tips').show();
            if (flag) {
                flag = false;
                sortableObj = Sortable.create(document.getElementById('sortable'), {
                    group: "dept",
                    sort: true,
                    animation: 150, //动画参数
                    onEnd: function (evt) { //拖拽完毕之后发生该事件
                        idList = [];
                        for (var i = 0; i < evt.from.children.length; i++) {
                            idList.push(evt.from.children[i].getAttribute('id'));
                        }
                    }
                });
            }
        })

        //保存排序
        $(document).off('click', '#saveDept');
        $(document).on('click', '#saveDept', function () {
            sortDestroy();
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/department-sort' //实际使用请改成服务端真实接口
                , type: 'put'
                , data: {'idList': idList, 'sortList': sortList}
                , done: function (res) {
                    var nodes = getZtreeObj(baseDeptId).getNodeByParam('name', deptName, null);
                    getDept(deptId, 'tree', nodes);//递归子部门
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        })

        //取消排序
        $(document).off('click', '#cancelDept');
        $(document).on('click', '#cancelDept', function () {
            sortDestroy();
            getDept(deptId, 'list');
        })

        //子部门操作
        $(document).off('click', '.sortable li .sort-btn,#deleteParent');
        $(document).on('click', '.sortable li .sort-btn,#deleteParent', function () {
            var deptId_ = '';
            if ($(this).hasClass('layui-btn')) {
                deptId_ = deptId;
            } else {
                deptId_ = $(this).parent().attr('id');
            }

            var type = $(this).data('type');
            if (type == 'name') {
                var deptName = $(this).text();
                getDept(deptId_, 'list');
                getMember(deptId_);
                saveDept(deptId_, deptName, baseDeptId);
                var param = getZtreeObj(baseDeptId).getNodeByParam('name', deptName, null);
                getZtreeObj(baseDeptId).selectNode(param);

                $('#name').html(deptName);
                var param = getZtreeObj(baseDeptId).getNodeByParam('name', deptName, null);
                if (!param.parentTId) {
                    $('#deleteParent').addClass('edit-show');
                } else {
                    $('#deleteParent').removeClass('edit-show');
                }
            } else {
                delDept(deptId_);
            }
        })

        //添加下级部门面板
        $(document).off('click', '#addDept');
        $(document).on('click', '#addDept', function () {
            endSort(function () {
                admin.popupCenter({
                    id: 'LAY_adminPopupAddDept'
                    , title: '添加下级部门'
                    , area: ['400px', '250px']
                    , success: function () {
                        view(this.id).render('setting/admin/organization/add_dept');
                    }
                });
            })
        })

        //添加一级部门面板
        $(document).off('click', '#addParentDept');
        $(document).on('click', '#addParentDept', function () {
            endSort(function () {
                admin.popupCenter({
                    id: 'LAY_adminPopupAddParentDept'
                    , title: '添加一级部门'
                    , area: ['400px', '200px']
                    , success: function () {
                        view(this.id).render('setting/admin/organization/addparent');
                    }
                });
            })
        })

        //添加成员面板
        $(document).off('click', '#addMember');
        $(document).on('click', '#addMember', function () {
            endSort(function () {
                admin.popupCenter({
                    id: 'LAY_adminPopupEditMember'
                    , title: '添加成员'
                    , area: ['400px', '420px']
                    , success: function () {
                        view(this.id).render('setting/admin/organization/addmember');
                    }
                });
            })
        })

        //成员权限管理面板
        $(document).off('click', '#authority');
        $(document).on('click', '#authority', function () {
            endSort(function () {
                var idList = getMemberId();
                layui.data('authorityUser', {
                    key: 'userList'
                    , value: idList
                });
                if (idList.length > 0) {
                    admin.popupRight({
                        id: 'LAY_adminPopupAuthority'
                        , title: '成员权限管理'
                        , area: '80%'
                        , closeBtn: 1
                        , success: function () {
                            view(this.id).render('setting/admin/organization/authority');
                        }
                    });
                } else {
                    layer.alert('请选择要操作的成员！', {icon: 5});
                }
            })
        })

        //调整部门
        $(document).off('click', '#modifyDept');
        $(document).on('click', '#modifyDept', function () {
            endSort(function () {
                if (getMemberId().length > 0) {
                    getMemberDept(getMemberId()[0]);
                } else if (getMemberId().length > 1) {
                    layui.data('temporary', {key: 'temporary', value: ''})
                    popupOrg('single', '调整部门', 'tree-box', 'modifydept');
                } else {
                    layer.alert('请选择要操作的成员！', {icon: 5});
                }
            })
        })

        function getMemberDept(id) {
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/member/' + id + '/departments'
                , type: 'get'
                , done: function (res) {
                    layui.data('temporary',
                        {
                            key: 'temporary',
                            value: res.data
                        }
                    )
                    layui.each(res.data,function(index,item){
                        deptIdList.push(item.id);
                        console.log(deptIdList)
                    })
                    popupOrg('single', '调整部门', 'tree-box', 'modifydept');
                }
            });
        }

        //导出成员
        $(document).off('click', '#export');
        $(document).on('click', '#export', function () {
            endSort(function () {
                popupOrg('more', '导出成员', 'export-tree', 'export', 'check');
            })
        })

        //导入成员
        $(document).off('click', '#import');
        $(document).on('click', '#import', function () {
            endSort(function () {
                admin.popupCenter({
                    id: 'LAY_adminPopupImportMember'
                    , title: '导入成员'
                    , area: ['600px', '370px']
                    , success: function () {
                        view(this.id).render('setting/admin/organization/importmember');
                    }
                });
            })
        })

        //编辑部门名称
        $(document).off('click', '#editname');
        $(document).on('click', '#editname', function () {
            endSort(function () {
                var content = '<div class="layui-form full-height centerLayer">' +
                    '<div class="layui-form-item">' +
                    '<label class="layui-form-label left">部门名称</label>' +
                    '<div class="layui-input-block">' +
                    '<input type="text" id="dept_name" autocomplete="off" value="' + layui.data('currentDept').dept.name + '" placeholder="请输入部门名称" class="layui-input">' +
                    '</div>' +
                    '</div>' +
                    '</div>';
                admin.popupCenter({
                    id: 'LAY_adminRolename'
                    , title: '修改部门名称'
                    , area: ['400px', '200px']
                    , btn: ['确定', '取消']
                    , content: content
                    , yes: function () {
                        var name = $('#dept_name').val();
                        if (!name) {
                            layer.msg('请输入模板名称', {icon: 5, time: 1000});
                        } else {
                            reset(deptId, name);
                        }
                    }
                });
            })
        })

        //批量删除成员
        $(document).off('click', '#batchDelete');
        $(document).on('click', '#batchDelete', function () {
            endSort(function () {
                var idList = getMemberId();
                if (idList.length > 0) {
                    layer.confirm('是否删除选中成员？', {icon: 3, title: '提示'}, function (index) {
                        layer.close(index);
                        admin.req({
                            url: setter.baseUrl + '/admin-manage/organize/member'
                            , type: 'delete'
                            , data: {'idList': idList, 'departmentId': deptId}
                            , done: function (res) {
                                getMember(deptId);
                                layer.msg(res.message, {icon: 1, time: 2000});
                            }
                        });
                    });
                } else {
                    layer.alert('请选择要操作的成员！', {icon: 5});
                }
            })
        })

        // 删除已选中部门（弹出层）
        $(document).off('click','#deleteDept');
        $(document).on('click','#deleteDept',function(){
            var this_elem = $(this);
            layer.confirm('是否删除该部门？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                deptIdList = layui.data('temporary').temporary;
                var id = this_elem.data('id');
                var idx = deptIdList.indexOf(id);
                deptIdList.splice(idx, 1);
                this_elem.parents('.deptname').remove();
                // setDownHref();
            });
        })

        //提交调整部门
        $(document).off('click', '#submitModify');
        $(document).on('click', '#submitModify', function () {
            var memberIdList = getMemberId();
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/member-department'
                , type: 'put'
                , data: {'departmentList': deptIdList, 'memberIdList': memberIdList}
                , done: function (res) {
                    getMember(deptId);
                    layer.close(admin.popup.index);
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        })

        //导出
        $(document).off('click', '#submit-export');
        $(document).on('click', '#submit-export', function () {
            $.each(zTreeObj_check, function (i, k) {
                var current_tree = zTreeObj_check[i].zTreeObj;
                $.each(current_tree.getCheckedNodes(true), function (i2, k2) {
                    deptIdList.push(k2.id);
                })
            })
            if (deptIdList.length == 0) {
                layer.alert('请先选择要导出的部门', {icon: 5})
            } else {
                setDownHref();
            }

        })

        //部门人员中搜索成员
        $(document).off('click', '#searchMember');
        $(document).on('click', '#searchMember', function () {
            endSort(function () {
                var queryString = $('#queryMember').val();
                getMember(deptId, queryString);
            })
        })
        $('#queryMember').bind('input propertychange', function () {
            var queryString = $(this).val();
            if (queryString.length == 0) {
                getMember(deptId);
            }
        })
        $('#queryMember').keypress(function (event) {
            var queryString = $(this).val();
            var keynum = (event.keyCode ? event.keyCode : event.which);
            if (keynum == '13') {
                getMember(deptId, queryString);
            }
        });
        //组织架构中搜索部门和成员
        $(document).off('click', '#globalSearch');
        $(document).on('click', '#globalSearch', function () {
            endSort(function () {
                var queryString = $('#globalKey').val();
                getResult(queryString);
            })
        })
        $('#globalKey').bind('input propertychange', function () {
            var queryString = $(this).val();
            if (queryString.length == 0) {
                $('.search-result').html('').hide();
            }
        })
        $('#globalKey').keypress(function (event) {
            var queryString = $(this).val();
            var keynum = (event.keyCode ? event.keyCode : event.which);
            if (keynum == '13') {
                getResult(queryString);
            }
        });

        //锁定搜索的部门
        $(document).off('click', '.result-dept');
        $(document).on('click', '.result-dept', function () {
            var userName = $(this).parent('.result-list').attr('data-userName');
            var deptId = $(this).attr('data-deptId');
            var deptName = $(this).attr('data-deptName');
            var baseDeptId = $(this).attr('data-baseDeptId');
            console.log(userName)
            console.log(deptId)
            console.log(deptName)
            console.log(baseDeptId)
            $('.search-result').hide();
            $('#globalKey').val('');
            $('#name').html(deptName);
            var param = getZtreeObj(baseDeptId).getNodeByParam('id', deptId, null);
            console.log(dept.baseDeptId);
            getZtreeObj(dept.baseDeptId).cancelSelectedNode();
            saveDept(deptId, deptName, baseDeptId);
            getDept(deptId, 'list');
            getZtreeObj(baseDeptId).selectNode(param);
            getMember(deptId, userName);
        })

        //点击节点操作
        function zTreeOnClickMain(event, treeId, treeNode) {
            if (!flag) {
                var param = getZtreeObj(baseDeptId).getNodeByParam('name', deptName, null);
                getZtreeObj(baseDeptId).selectNode(param);
                getZtreeObj(treeId.replace(/[^0-9]/ig, "")).cancelSelectedNode();
            }
            endSort(function () {
                $('.option-tips').hide();
                var currentId = treeId.replace(/[^0-9]/ig, "");
                for (i = 0; i < zTreeObj.length; i++) {
                    var id = zTreeObj[i].baseDeptId;
                    if (id != currentId) {
                        getZtreeObj(id).cancelSelectedNode();
                    }
                }
                baseDeptId = currentId;
                saveDept(treeNode.id, treeNode.name, currentId);

                $('#name').html(treeNode.name);
                var param = getZtreeObj(baseDeptId).getNodeByParam('name', deptName, null);
                if (!param.parentTId) {
                    $('#deleteParent').addClass('edit-show');
                } else {
                    $('#deleteParent').removeClass('edit-show');
                }

                getDept(treeNode.id, 'list');
                getMember(treeNode.id);
            })
        }

        //点击节点操作(弹出层)
        function zTreeOnClick(event, treeId, treeNode) {
            // if (treeSelect == 'single') {
            deptIdList = layui.data('temporary').temporary;
            console.log(deptIdList)
            $.each(zTreeObj_, function (i, k) {
                zTreeObj_[i].zTreeObj.cancelSelectedNode();
                zTreeObj_[i].zTreeObj.selectNode(treeNode);
            })
            departmentId = treeNode.id;
            if(deptIdList.indexOf(departmentId)<0){
                deptIdList.push(departmentId);
                $('#selectedDept').append('<p class="deptname">' + treeNode.name +
                    '<svg class="icon pull-right" aria-hidden="true" id="deleteDept" data-id="' + treeNode.id + '"><use xlink:href="#icon-delete"></use></svg>' +
                    '</p>');
            }else{
                layer.alert('该部门已被选中', {icon: 5});
            }

            // }
            // else {
            //     if (deptIdList.indexOf(treeNode.id) == -1) {
            //         $('#selectedDept').append('<p class="deptname">' + treeNode.name +
            //             '<svg class="icon pull-right" aria-hidden="true" id="deleteDept" data-id="' + treeNode.id + '"><use xlink:href="#icon-delete"></use></svg>' +
            //             '</p>');
            //         deptIdList.push(treeNode.id);
            //         setDownHref();
            //     }
            // }
        }

        //调用组织架构弹窗
        function popupOrg(type, title, elem, html_name, treetype) {
            admin.popupCenter({
                id: 'LAY_adminPopupModifyDept'
                , title: title
                , area: ['640px', '600px']
                , success: function () {
                    view(this.id).render('setting/admin/organization/' + html_name);
                    getTree(elem, treetype);
                }
            });
            treeSelect = type;
        }

        //获取组织架构
        var count = 0;

        function getTree(elem, type) {    //id：所在部门id
            $('.' + elem).html('');
            var roleIdArr = [];
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/department' //实际使用请改成服务端真实接口
                , done: function (res) {
                    var html = '';
                    $.each(res.data, function (i, k) {
                        html = '<ul class="ztree" id="' + elem + k.id + '"></ul>';

                        $('.' + elem).append(html);
                        //zTree配置
                        var zNodes = [k];
                        if (type == 'initPage') {
                            var setting = {
                                callback: {
                                    onClick: zTreeOnClickMain
                                }
                            };
                            var zTreeObj1 = $.fn.zTree.init($('#' + elem + k.id), setting, zNodes);
                            zTreeObj.push({'baseDeptId': k.id, 'zTreeObj': zTreeObj1});
                        } else if (type == 'check') {
                            var setting = {
                                check: {
                                    enable: true,
                                    chkboxType: {"Y": "ps", "N": "s"}
                                }
                            };
                            var zTreeObj1 = $.fn.zTree.init($('#' + elem + k.id), setting, zNodes);
                            zTreeObj_check.push({'baseDeptId': k.id, 'zTreeObj': zTreeObj1});
                        } else {
                            var setting = {
                                callback: {
                                    onClick: zTreeOnClick
                                }
                            };
                            var zTreeObj1 = $.fn.zTree.init($('#' + elem + k.id), setting, zNodes);
                            zTreeObj_.push({'baseDeptId': k.id, 'zTreeObj': zTreeObj1});
                        }

                    })

                    var k = res.data[0];
                    if (res.data.length > 0 && count == 0 && type == 'initPage') {
                        deptId = k.id;
                        deptName = k.name;
                        baseDeptId = k.id;
                        saveDept(deptId, deptName, baseDeptId);//存储当前部门信息
                        //zTree选中
                        var param = getZtreeObj(baseDeptId).getNodeByParam('name', deptName, null);
                        getZtreeObj(baseDeptId).selectNode(param);

                        $('#name').html(deptName);
                        getDept(deptId, 'list');//获取直接子部门
                        getMember(deptId);//获取成员信息
                        if (!param.parentTId) {
                            $('#deleteParent').addClass('edit-show');
                        } else {
                            $('#deleteParent').removeClass('edit-show');
                        }
                        count++;
                    }
                }
            });
        }

        //获取部门成员信息
        function getMember(id, other) {    //id：所在部门id  other:其他搜索字段
            layer.closeAll('tips');
            if (other == undefined) {
                other = '';
            }
            table.render({
                id: 'member'
                , elem: '#member'
                , url: setter.baseUrl + '/admin-manage/organize/department-member' //模拟接口
                , page: setter.pageTable
                , cellMinWidth: 80
                , loading: false
                , limit: setter.limit
                , response: setter.responseTable
                , cols: [[
                    {type: 'checkbox', fixed: 'left'}
                    , {type: 'numbers', title: '序号'}
                    , {field: 'name', title: '姓名', width: 100, templet: '#is_admin'}
                    , {field: 'mobile', title: '手机号', width: 150}
                    , {field: 'job_number', title: '工号'}
                    , {field: 'post_name', title: '岗位'}
                    , {field: 'email', width: 150, title: '邮箱'}
                    , {title: '操作', width: 150, toolbar: '#option'}
                    // ,{field: 'isTop', title:'置顶', width:85, templet: '#switchTpl', unresize: true}
                ]]
                , where: {
                    'id': id
                    , 'queryString': other
                }
                , skin: 'line'
            });
        }

        //监听置顶操作
        var current_member_id = '';
        table.on('tool(memberTable)', function (obj) {
            var data = obj.data;
            if (obj.event === 'toTop') {
                admin.req({
                    url: setter.baseUrl + '/admin-manage/organize/member-top'
                    , type: 'put'
                    , data: {'id': data.id}
                    , done: function (res) {
                        getMember(deptId);
                    }
                });
            } else {
                layui.data('currentMember', {
                    key: 'mobile'
                    , value: data.mobile
                });
                current_member_id = data.id;
                endSort(function () {
                    admin.popupCenter({
                        id: 'LAY_adminPopupEditMember'
                        , title: '编辑成员'
                        , area: ['400px', '420px']
                        , success: function () {
                            view(this.id).render('setting/admin/organization/editmember');
                        }
                    });
                })
            }
        });

        //获取子部门列表
        function getDept(id, type, nodes) {    //id：上级部门id  type: list直接子部门，tree递归子部门  //nodes: 父类树节点，用于更新树
            var html = '';
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/department-child' //实际使用请改成服务端真实接口
                , data: {'id': id, 'type': type}
                , done: function (res) {
                    if (type == 'list') {
                        sortList = [];
                        idList = [];
                        if (res.data.length > 0) {
                            $('#sort').addClass('sort');
                            $.each(res.data, function (i, k) {
                                sortList.push(k.sort);
                                idList.push(k.id);
                                html += '<li class="ui-state-default layui-clear" id="' + k.id + '">' +
                                    '<span class="layui-col-sm11 sort-btn" data-type="name">' + k.name + '</span>' +
                                    '<a href="javascript:;" class="sort-btn ft-right layui-col-sm1" data-type="delete">' +
                                    '<svg class="icon" aria-hidden="true">' +
                                    '<use xlink:href="#icon-delete"></use>' +
                                    '</svg>' +
                                    '</a></li>';
                            })
                        } else {
                            $('#sort').removeClass('sort');
                            html = '<p style="padding: 20px;text-align: center;">当前部门无下级部门，' +
                                '<a href="javascript:;" id="addDept" style="color: #1E9FFF;">请添加子部门</a></p>';
                        }
                        $('#sortable').html(html);
                    } else {
                        nameList = res.data;
                        $('.option-tips').hide();
                        // console.log(nodes)
                        // console.log(nameList)
                        getZtreeObj(baseDeptId).removeChildNodes(nodes);
                        getZtreeObj(baseDeptId).addNodes(nodes, nameList, null);
                        var param = getZtreeObj(baseDeptId).getNodeByParam('name', deptName, null);
                        getZtreeObj(baseDeptId).selectNode(param);
                    }
                }
            });
        }

        //数据存储
        function saveDept(id, name, baseDeptId) {
            layui.data('currentDept', {
                key: 'dept'
                , value: {'id': id, 'name': name, 'baseDeptId': baseDeptId}
            });
            dept = layui.data('currentDept').dept;
            deptId = dept.id;
            deptName = dept.name;
            baseDeptId = dept.baseDeptId;
        }

        //删除部门
        function delDept(id) {
            layer.confirm('是否删除该部门？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + '/admin-manage/organize/department/' + id
                    , type: 'delete'
                    , done: function (res) {
                        var param = getZtreeObj(baseDeptId).getNodeByParam('id', id, null);
                        if (!param.parentTId) {
                            count = 0;
                            getTree('result', 'initPage')
                        } else {
                            getDept(deptId, 'list');
                            getZtreeObj(baseDeptId).removeNode(param);
                        }

                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                });
            });
        }

        //获取表格中选中成员id
        function getMemberId() {
            var memberId = [];
            var checkStatus = table.checkStatus('member'), data = checkStatus.data;
            for (var i in data) {
                memberId.push(data[i].id);
            }
            return memberId;
        }

        //提交编辑部门名称
        function reset(id, name) {
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/department-name'
                , type: 'put'
                , data: {'id': id, 'departmentName': name}
                , done: function (res) {
                    var nodes = getZtreeObj(baseDeptId).getNodesByParam("name", deptName, null);
                    if (!nodes[0].parent_id) {
                        count = 0;
                        getTree('result', 'initPage');
                    } else {
                        $('#name').html(name);
                        var parentNode = getZtreeObj(baseDeptId).getNodeByTId(nodes[0].parentTId);
                        var parentId = parentNode.id;
                        getDept(parentId, 'tree', parentNode);
                        saveDept(deptId, name, baseDeptId);
                    }
                    layer.msg(res.message, {icon: 1, time: 2000});
                    layer.close(admin.popup.index);//关闭面板
                }
            });
        }

        //获取组织架构搜索结果
        function getResult(queryString) {
            var html = '',list = '';
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/department-or-member'
                , type: 'get'
                , data: {'queryString': queryString}
                , done: function (res) {
                    if (res.data.admin.length > 0) {
                        $.each(res.data.admin, function (i, k) {
                            list = '';
                            $.each(k.departments, function (i2, k2) {
                                list += '<p class="result-dept" data-deptId="' + k2.id + '" data-deptName="' + k2.name + '" data-baseDeptId="' + k2.baseDepartId + '">' + k2.name + '<span class="layui-icon layui-icon-right pull-right"></span></p>';
                            })
                            html += '<li class="result-list" data-userName="' + k.name + '"><p class="ft-green">'+ k.name +'<span class="pull-right">'+ k.mobile +'</span></p>'+ list +'</li>'
                        })
                    } else if (res.data.department.length > 0) {
                        $.each(res.data.department, function (i, k) {
                            html += '<li class="result-list"><p class="result-dept" data-deptId="' + k.id + '" data-userName="" data-deptName="' + k.name + '" data-baseDeptId="' + k.baseDepartId + '">' + k.name + '</p></li>';
                        })
                    } else {
                        html += '<li class="result-list">没有符合条件的结果</li>';
                    }
                    $('.search-result').html(html).show();
                }
            });
        }

        //生成导出地址
        function setDownHref() {
            downHref = setter.baseUrl + '/admin-manage/organize/export/member-excel?';
            $.each(deptIdList, function (i, k) {
                downHref += 'departmentList[]=' + k + '&';
            })
            // $('#submit-export').attr('href',downHref);
            window.open(downHref);
            layer.close(admin.popup.index);
        }

        //获取树对象
        function getZtreeObj(baseDeptId) {
            var param;
            $.each(zTreeObj, function (i, k) {
                if (k.baseDeptId == baseDeptId) {
                    param = zTreeObj[i].zTreeObj;
                }
            })
            return param;
        }

        //销毁sortable
        function sortDestroy() {
            sortableObj.destroy();
            $('.option-tips').hide();
            $('#sortable').addClass('sortable');
            flag = true;
        }

        //提示先结束排序
        function endSort(callback) {
            if (flag) {
                callback();
            } else {
                layer.alert('请先结束排序操作！', {icon: 5});
            }
        }

        /*addDept*/
        //添加部门
        form.on('submit(LAY-add-dept-submit)', function (obj) {
            //请求添加部门接口
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/department'
                , type: 'post'
                , data: obj.field
                , done: function (res) {
                    //关闭右侧面板
                    layer.close(admin.popup.index);
                    if (obj.field.id == 0) {
                        count = 0;
                        getTree('result', 'initPage');
                    } else {
                        var nodes = getZtreeObj(baseDeptId).getNodeByParam('name', deptName, null);
                        getDept(deptId, 'list');
                        getDept(deptId, 'tree', nodes);
                    }
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        });

        /*addMember*/
        //添加成员
        form.on('submit(LAY-add-member-submit)', function (obj) {
            //请求添加部门接口
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/department-member'
                , type: 'post'
                , data: obj.field
                , done: function (res) {
                    //关闭右侧面板
                    layer.close(admin.popup.index);
                    getMember(deptId);
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        });

        /*editMember*/
        //编辑成员
        form.on('submit(LAY-edit-member-submit)', function (obj) {
            //请求添加部门接口
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/department-member/' + current_member_id
                , type: 'put'
                , data: obj.field
                , done: function (res) {
                    //关闭右侧面板
                    layer.close(admin.popup.index);
                    getMember(deptId);
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        });

        /*authority*/
        //权限管理
        $(document).off('click', '#authorityConfirm');
        $(document).on('click', '#authorityConfirm', function () {
            var userList = layui.data('authorityUser').userList;
            var permissionList = [];
            $("input:checkbox[name='authority']:checked").each(function () {
                permissionList.push($(this).val());
            });
            admin.req({
                url: setter.baseUrl + '/admin-manage/organize/assign-permission'
                , type: 'post'
                , data: {'userList': userList, 'permissionList': permissionList}
                , done: function (res) {
                    layer.close(admin.popup.index);
                    layui.data('authorityUser', null);
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        })

        //关闭弹窗
        $(document).off('click', '#cancel');
        $(document).on('click', '#cancel', function () {
            layer.close(admin.popup.index);
        })

        //确认上传文件
        $(document).off('click', '#importUser');
        $(document).on('click', '#importUser', function () {
            var excelPath = $('#excelPath').val();
            if (!excelPath) {
                layer.alert('请先上传需要导入的文件！', {icon: 5});
            } else {
                admin.req({
                    url: setter.baseUrl + '/admin-manage/organize/import/member-excel' //实际使用请改成服务端真实接口
                    , type: 'post'
                    , data: {'excelPath': excelPath}
                    , done: function (res) {
                        count = 0;
                        layer.close(admin.popup.index);
                        layer.msg(res.message, {icon:1, time: 2000});
                        getTree('result', 'initPage');
                    }
                });
            }
        })

        $(document).click(function () {
            $('.search-result').html('').hide();
        })
    });
    exports('setting/organization', {})
});