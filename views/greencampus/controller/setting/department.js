layui.define(['admin', 'table'], function (exports) {
    var $ = layui.$,
        admin = layui.admin,
        table = layui.table,
        setter = layui.setter,
        view = layui.view,
        form = layui.form;

    $(function () {
        // var line_width = $('.content-line').outerWidth();

        var id,name;

        init();

        function init() {
            admin.req({
                url: setter.baseUrl + '/user-manage/role/' + layui.data('roleParam').id + '/structure-init'
                , type: 'get'
                , done: function (res) {
                    // view('modal-content').refresh(function () {
                    //     element.tabChange('tab-child', id);
                    // });
                    // layer.msg(res.message, {icon: 1, time: 2000});
                    var html = '', parentId = '';
                    $.each(res.initStructures, function (index, item) {
                        var content = '';
                        $.each(item.structures, function (index2, item2) {
                            var active = '';
                            if (index2 === 0) {
                                active = 'active'
                            } else {
                                active = '';
                            }
                            content += '<li class="dept-item ' + active + ' flex flex-align-center dept-item-'+ item2.id +'" data-id="' + item2.id + '" data-name="' + item2.name + '">' +
                                '<span class="text flex-1">' + item2.name + ' </span>' +
                                '<svg class="icon goto" aria-hidden="true"><use xlink:href="#icon-fh"></use></svg>' +
                                '<div class="icon-btn-box">' +
                                '<svg class="icon" id="icon-bianji" aria-hidden="true" data-id="'+ item2.id +'" data-name="'+ item2.name +'"><use xlink:href="#icon-bianji"></use></svg>' +
                                '<svg class="icon" id="icon-delete" aria-hidden="true" data-id="'+ item2.id +'"><use xlink:href="#icon-delete"></use></svg>' +
                                '</div>' +
                                '</li>';
                        })
                        html += '<div class="content-line full-height">' +
                            '<div class="layui-card full-height flex flex-v">' +
                            '<div class="layui-card-header">' + item.title + ' </div>' +
                            '<div class="layui-card-body flex-1 over-auto">' +
                            '<div class="layui-form-item hide">' +
                            '<input type="text" name="name" lay-verify="required" autocomplete="off" value="" placeholder="请输入名称" class="layui-input">' +
                            '</div>' +
                            '<ul class="dept-list">' + content + '</ul>' +
                            '</div>' +
                            '<div class="layui-card-body flex">' +
                            '<button class="layui-btn layui-btn-primary flex-1 addStructure">' +
                            '<svg class="icon" aria-hidden="true"><use xlink:href="#icon-jia"></use></svg>新增' +
                            '</button>' +
                            '</div>' +
                            '</div>' +
                            '</div>'
                    })

                    $('#structures').html(html);
                    var length = $('.content-line').length;
                    $('.content').css('width', 300 * length);
                }
            });
        }

        var parentId,currentAdd;
        $(document).off('click', '.addStructure');
        $(document).on('click', '.addStructure', function () {
            currentAdd = $(this);
            parentId = $(this).parents('.content-line').prev().find('.active').attr('data-id');
            if (!parentId) {
                parentId = 0;
            }
            console.log(parentId)
            layui.data('temporary', {
                key: 'temporary',
                value: {parentId: parentId}
            });
            admin.popupCenter({
                id: 'LAY_adminPopupAddrRoleDept'
                , title: '添加选项'
                , area: ['400px', '200px']
                , success: function () {
                    view(this.id).render('setting/user/role/addDept');
                }
            });
        })


        //添加
        form.on('submit(LAY-add-role-dept)', function (obj) {
            obj.field.roleId = layui.data('roleParam').id;
            obj.field.parentId = parentId;
            admin.req({
                url: setter.baseUrl + '/user-manage/role/structure'
                , type: 'post'
                , data: obj.field
                , done: function (res) {
                    //关闭右侧面板
                    layer.close(admin.popup.index);
                    var content = '<li class="dept-item flex flex-align-center dept-item-'+ res.data.id +'" data-id="' + res.data.id + '" data-name="' + obj.field.name + '">' +
                        '<span class="text flex-1">' + obj.field.name + ' </span>' +
                        '<svg class="icon goto" aria-hidden="true"><use xlink:href="#icon-fh"></use></svg>' +
                        '<div class="icon-btn-box">' +
                        '<svg class="icon" id="icon-bianji" aria-hidden="true" data-id="'+ res.data.id +'" data-name="'+ obj.field.name +'"><use xlink:href="#icon-bianji"></use></svg>' +
                        '<svg class="icon" id="icon-delete" aria-hidden="true" data-id="'+ res.data.id +'"><use xlink:href="#icon-delete"></use></svg>' +
                        '</div>' +
                        '</li>';
                    currentAdd.parent().siblings().find('.dept-list').append(content)
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        });

        //修改
        form.on('submit(LAY-reset-role-dept)', function (obj) {
            admin.req({
                url: setter.baseUrl + '/user-manage/role/structure/'+ id
                , type: 'put'
                , data: obj.field
                , done: function (res) {
                    //关闭右侧面板
                    layer.close(admin.popup.index);
                    $('.dept-item-'+ id).find('.text').html(obj.field.name);
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        });

        $(document).off('click', '.dept-item');
        $(document).on('click', '.dept-item', function () {
            id = $(this).attr('data-id');
            name = $(this).attr('data-name');
            $(this).addClass('active').siblings().removeClass('active');
            var elem = $(this).parents('.content-line').nextAll();
            admin.req({
                url: setter.baseUrl + '/user-manage/role/structure/' + id + '/sons'
                , type: 'get'
                , done: function (res) {
                    var contentArr = [];
                    if (JSON.stringify(res.data) != "{}") {
                        $.each(res.data, function (index, item) {
                            var content = '';
                            $.each(item, function (index2, item2) {
                                var active = '';
                                if (index2 === 0) {
                                    active = 'active'
                                } else {
                                    active = '';
                                }
                                content += '<li class="dept-item ' + active + ' flex flex-align-center dept-item-'+ item2.id +'" data-id="' + item2.id + '">' +
                                    '<span class="text flex-1">' + item2.name + ' </span>' +
                                    '<svg class="icon goto" aria-hidden="true"><use xlink:href="#icon-fh"></use></svg>' +
                                    '<div class="icon-btn-box">' +
                                    '<svg class="icon" id="icon-bianji" aria-hidden="true" data-id="'+ item2.id +'" data-name="'+ item2.name +'"><use xlink:href="#icon-bianji"></use></svg>' +
                                    '<svg class="icon" id="icon-delete" aria-hidden="true" data-id="'+ item2.id +'"><use xlink:href="#icon-delete"></use></svg>' +
                                    '</div>' +
                                    '</li>';
                            });
                            contentArr.push(content);
                        })
                        if (contentArr.length > 0) {
                            elem.find('.dept-list').html('');
                            elem.each(function (i, k) {
                                $(this).find('.dept-list').html(contentArr[i]);
                            })
                        } else {
                            elem.find('.dept-list').html('');
                        }
                    }
                }
            })
        });

        // 编辑
        $(document).off('click', '#icon-bianji');
        $(document).on('click', '#icon-bianji', function (e) {
            e.stopPropagation();
            id = $(this).attr('data-id');
            name = $(this).attr('data-name');
            layui.data('temporary', {
                key: 'temporary',
                value: {id: id, name: name}
            });
            admin.popupCenter({
                id: 'LAY_adminPopupResetrRoleDept'
                , title: '添加选项'
                , area: ['400px', '200px']
                , success: function () {
                    view(this.id).render('setting/user/role/resetDept');
                }
            });
        })

        // 删除
        $(document).off('click', '#icon-delete');
        $(document).on('click', '#icon-delete', function (e) {
            e.stopPropagation();
            var id = $(this).attr('data-id');
            layer.confirm('是否删除该选项？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + '/user-manage/role/structure/'+id
                    , type: 'delete'
                    , data: {}
                    , done: function (res) {
                        //关闭右侧面板
                        layer.close(admin.popup.index);
                        $('.dept-item-'+ id).remove();
                        layer.msg(res.message, {icon: 1, time: 2000});
                    }
                });
            });

        })

        //关闭弹窗
        $(document).off('click', '#cancel');
        $(document).on('click', '#cancel', function () {
            layer.close(admin.popup.index);
        })

    });
    exports('setting/department', {})
});