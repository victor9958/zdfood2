layui.define(['admin', 'form', 'zTree'], function (exports) {
    var $ = layui.$,
        admin = layui.admin,
        setter = layui.setter,
        form = layui.form,
        view = layui.view;

    var tree, nodes, parentId;
    $(function () {
        var setting = {
            view: {
                expandSpeed: "",
                addHoverDom: addHoverDom,
                removeHoverDom: removeHoverDom,
                selectedMulti: false
            },
            async: {
                enable: true,
                type: 'get',
                url: setter.baseUrl + '/system-set/site/child',
                autoParam: ["id"]
            },
            edit: {
                drag: {
                    autoExpandTrigger: true,
                    prev: true,
                    inner: true,
                    next: true
                },
                enable: true,
                showRemoveBtn: true,
                showRenameBtn: true
            },
            callback: {
                onClick: zTreeOnClick,
                onAsyncError: onAsyncError,
                onAsyncSuccess: onAsyncSuccess,
                // onDrop: onDrop,
                beforeDrop: beforeDrop,
                beforeRemove: beforeRemove,
                beforeRename: beforeRename,
                onRename: onRename
            }
        };

        function beforeRemove(treeId, treeNode) {
            var idList = [treeNode.id];
            // console.log(treeNode);
            tree.selectNode(treeNode);
            layer.confirm('是否删除所有选中的数据？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + '/system-set/sites'
                    , type: 'delete'
                    , data: {'idList': idList}
                    , done: function (res) {
                        var parentNode;
                        tree.removeNode(treeNode);
                        if (treeNode.parentTId != null) {
                            parentNode = tree.getNodeByTId(treeNode.parentTId);
                            parentId = parentNode.id;
                        } else {
                            parentId = 0;
                            parentNode = tree.getNodes();
                        }
                        tree.selectNode(parentNode);
                        var name = parentNode.name;
                        $('#levelname').html(name);
                        getChild();
                        layer.close(admin.popup.index);//关闭右侧面板
                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                });
            }, function () {
                layer.close(admin.popup.index);//关闭右侧面板
            });
            return false;
        }

        var new_name;

        function beforeRename(treeId, treeNode, newName) {
            if (newName.length == 0) {
                layer.alert("节点名称不能为空", {icon: 5}, function (index) {
                    layer.close(index);
                    $('.rename').focus();
                });
                return false;
            } else {
                new_name = newName;
                parentId = treeNode.id;
                tree.selectNode(treeNode);
                getChild();
            }
            return true;
        }

        //修改名称
        function onRename(event, treeId, treeNode, isCancel) {
            admin.req({
                url: setter.baseUrl + '/system-set/site/' + parentId
                , type: 'put'
                , data: {'name': new_name}
                , done: function (res) {
                    $('#levelname').html(new_name);
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        }

        function addHoverDom(treeId, treeNode) {
            var sObj = $("#" + treeNode.tId + "_span");
            if (treeNode.editNameFlag || $("#addBtn_" + treeNode.tId).length > 0) return;
            var addStr = "<span class='button add' id='addBtn_" + treeNode.tId + "' title='add node' onfocus='this.blur();'></span>";

            sObj.after(addStr);
            $("#addBtn_" + treeNode.tId).bind("click", function () {
                tree.selectNode(treeNode);
                parentId = treeNode.id;
                admin.popupCenter({
                    id: 'LAY_adminPopupAddChild'
                    , title: '添加子级信息'
                    , area: ['400px', '200px']
                    , success: function () {
                        view(this.id).render('setting/site/park/addchild');
                    }
                });
                return false;
            });
        };

        function removeHoverDom(treeId, treeNode) {
            $("#addBtn_" + treeNode.tId).unbind().remove();
        };

        function beforeDrop(treeId, treeNodes, targetNode, moveType, isCopy) {
            var msg = '是否将 [' + treeNodes[0].name + '] 移动到 [' + targetNode.name;
            switch (moveType) {
                case 'inner':
                    msg += ' ]之内?';
                    break;
                case 'prev':
                    msg += ' ]之前?';
                    break;
                case 'next':
                    msg += ' ]之后?';
                default:
                    break;
            }
            layer.confirm(msg, function (index) {
                tree.moveNode(targetNode, treeNodes[0], moveType);
                layer.close(index);

                var targetList = [], parentNode;
                var currentId = treeNodes[0].id;
                if (treeNodes[0].parentTId != null) {
                    parentNode = tree.getNodeByTId(treeNodes[0].parentTId);
                    if (parentNode.isParent) {
                        $.each(parentNode.children, function (i, k) {
                            targetList.push(k.id);
                        })
                        parentId = parentNode.id;
                    }
                } else {
                    parentId = 0;
                    parentNode = tree.getNodes();
                    $.each(parentNode, function (i, k) {
                        targetList.push(k.id);
                    })
                }

                admin.req({
                    url: setter.baseUrl + '/ystem-set/site/drag'
                    , type: 'put'
                    , data: {'id': currentId, 'parent_id': parentId, 'targetList': targetList}
                    , done: function (res) {
                        var name = treeNodes[0].name;
                        $('#levelname').html(name);
                        getChild();
                        // if(treeNodes[0].parentTId != null){
                        //     refreshNode('refresh', false, parentNode);
                        // }else{
                        //     tree = $.fn.zTree.init($("#parkTree"), setting);
                        // }
                        // tree = $.fn.zTree.init($("#parkTree"), setting);
                        // refreshNode('refresh', false ,parentNode);
                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                });
            });
            return false;
        }

        function onAsyncError(event, treeId, treeNode, XMLHttpRequest, textStatus, errorThrown) {
        }

        function onAsyncSuccess(event, treeId, treeNode, msg) {
            if(treeNode == undefined){

            }else if (!treeNode) {
                var name = tree.getNodes()[0].name;
                parentId = tree.getNodes()[0].id;
                $('#levelname').html(name);
                tree.selectNode(tree.getNodes()[0]);
            } else {
                $('#levelname').html(treeNode.name);
                parentId = treeNode.id;
                tree.selectNode(treeNode);
            }
            getChild();
        }

        function zTreeOnClick(event, treeId, treeNode) {
            parentId = treeNode.id;
            var name = treeNode.name;
            $('#levelname').html(name);
            getChild();
        }

        function refreshNode(type, silent, node) {
            if (!node) {
                nodes = tree.getSelectedNodes();
                tree.reAsyncChildNodes(nodes[0], type, silent);
            } else {
                nodes = node;
                tree.reAsyncChildNodes(nodes, type, silent);
            }
            if (!silent) tree.selectNode(nodes[0]);
        }

        function getChild() {
            admin.req({
                url: setter.baseUrl + '/system-set/site/child'
                , type: 'get'
                , data: {'id': parentId}
                , done: function (res) {
                    var html = '<li class="ui-state-default layui-form-item" id="" data-name="">' +
                        '<input type="checkbox" name="checkall" value="" title="全选" lay-skin="primary" lay-filter="allcheck">' +
                        '</li>';
                    if (res.length > 0) {
                        $('#parkSort').addClass('sort');
                        $.each(res, function (i, k) {
                            html += '<li class="ui-state-default layui-form-item check_child" id="' + k.id + '" data-name="' + k.name + '">' +
                                '<input type="checkbox" name="idList[]" value="' + k.id + '" title="' + k.name + '" lay-skin="primary" lay-filter="singlecheck">' +
                                '</li>';
                        })
                    } else {
                        $('#parkSort').removeClass('sort');
                        html = '<p style="padding: 20px;text-align: center;">当前层级无下级</p>';
                    }

                    $('#parkSortable').html(html);
                    form.render();
                }
            });
        }


        tree = $.fn.zTree.init($("#parkTree"), setting);

        //导入
        $(document).off('click', '#import_park');
        $(document).on('click', '#import_park', function () {
            admin.popupCenter({
                id: 'LAY_adminPopupImportPark'
                , title: '导入Excel数据'
                , area: ['600px', '550px']
                , success: function () {
                    view(this.id).render('setting/site/park/import');
                }
            });
        })

        //添加子级
        form.on('submit(LAY-add-child)', function (obj) {
            //请求添加部门接口
            obj.field.parentId = parentId;
            var treeNode = tree.getNodeByParam('id', parentId, null);
            admin.req({
                url: setter.baseUrl + '/system-set/site'
                , type: 'post'
                , data: obj.field
                , done: function (res) {
                    if(obj.field.id == 0){
                        layer.close(admin.popup.index);
                        layer.msg(res.message, {icon: 1, time: 2000});
                        tree = $.fn.zTree.init($("#parkTree"), setting);
                    }else{
                        tree.addNodes(treeNode, res.data);
                        getChild();
                        layer.close(admin.popup.index);
                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                }
            });
        });

        //添加一级面板
        $(document).off('click', '#addParent');
        $(document).on('click', '#addParent', function () {
            admin.popupCenter({
                id: 'LAY_adminPopupAddParent'
                , title: '添加一级'
                , area: ['400px', '200px']
                , success: function () {
                    parentId = 0;
                    view(this.id).render('setting/site/park/addchild');
                }
            });
        })

        //删除子级
        form.on('submit(LAY-delete-child)', function (obj) {
            var check_length = $('.check_child .layui-form-checked').length;//选中数量
            var all_length = $('.check_child .layui-form-checkbox').length;//所有数量
            var parentNode = tree.getNodeByParam('id', parentId, null);
            if ($.isEmptyObject(obj.field)) {
                layer.alert('请先选择要删除的记录', {icon: 5})
            } else {
                layer.confirm('是否删除所有选中的数据？', {icon: 3, title: '提示'}, function (index) {
                    layer.close(index);
                    admin.req({
                        url: setter.baseUrl + '/system-set/sites'
                        , type: 'delete'
                        , data: obj.field
                        , done: function (res) {
                            getChild();
                            if (check_length == all_length) {
                                tree.removeChildNodes(parentNode);
                            } else {
                                refreshNode('refresh', false);
                            }
                            layer.close(admin.popup.index);//关闭右侧面板
                            layer.msg(res.message, {icon: 1, time: 2000});
                        }
                    });
                });
            }
        });

        //全选
        form.on('checkbox(allcheck)', function (data) {
            var flag = data.elem.checked;
            if (flag) {
                $('#parkSortable').find('input').prop("checked", true);
                form.render('checkbox')
            } else {
                $('#parkSortable').find('input').prop("checked", false);
                form.render('checkbox')
            }
        });

        //单选
        form.on('checkbox(singlecheck)', function (data) {
            var flag = data.elem.checked;
            var check_length = $('.check_child .layui-form-checked').length;//选中数量
            var all_length = $('.check_child .layui-form-checkbox').length;//所有数量
            if (check_length == all_length) {
                $("input:checkbox[name='checkall']").prop("checked", true);
                form.render('checkbox');
            } else {
                $("input:checkbox[name='checkall']").prop("checked", false);
                form.render('checkbox');
            }
        });

        //关闭弹窗
        $(document).off('click', '#cancel');
        $(document).on('click', '#cancel', function () {
            layer.close(admin.popup.index);
        })

        //确认上传文件
        $(document).off('click', '#confirm_park');
        $(document).on('click', '#confirm_park', function () {
            var excelPath = $('#excelPath_park').val();
            if (!excelPath) {
                layer.alert('请先上传需要导入的文件！', {icon: 5});
            } else {
                admin.req({
                    url: setter.baseUrl + '/system-set/site/import-excel' //实际使用请改成服务端真实接口
                    , type: 'post'
                    , data: {'excelPath': excelPath}
                    , done: function (res) {
                        layer.close(admin.popup.index);
                        parentId = 0;
                        tree = $.fn.zTree.init($("#parkTree"), setting);
                    }
                });
            }
        })

    });
    exports('setting/park', {})
});