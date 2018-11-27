layui.define(['admin', 'upload', 'table', 'form'], function (exports) {
    var $ = layui.$,
        table = layui.table,
        admin = layui.admin,
        setter = layui.setter,
        form = layui.form;
    layui.admin.member = {
        tableData: function (tableName, url, query, col) {
            layer.closeAll('tips');
            table.render({
                id: tableName
                , elem: '#' + tableName
                , url: setter.baseUrl + url //模拟接口
                , page: setter.pageTable
                , cellMinWidth: 100
                , loading: false
                , limit: setter.limit
                , response: setter.responseTable
                , cols: [
                    col
                ]
                , where: query
                , skin: 'line'
            });
        },
        memberExport: function () {
            form.on('submit(LAY-exportMember)', function (obj) {
                var href = setter.baseUrl + '/member-manage/index/user-export?Authorization='
                    + layui.data('layuiAdmin').Authorization + '&operatorId='
                    + obj.field.operatorId + '&queryString=' + obj.field.queryString;
                window.location.href = href;
            });
        },
        rechargeExport: function () {
            form.on('submit(LAY-exportMember)', function (obj) {
                var href = setter.baseUrl + '/member-manage/recharge/records-export?Authorization='
                    + layui.data('layuiAdmin').Authorization + '&operatorId='
                    + obj.field.operatorId + '&type=' + obj.field.type + '&queryString=' + obj.field.queryString;
                window.location.href = href;
            });
        },
        cancel: function () {
            $(document).off('click', '#cancel');
            $(document).on('click', '#cancel', function () {
                layer.close(admin.popup.index);
            })
        },
    };

    exports('member/member', layui.admin.member)
})