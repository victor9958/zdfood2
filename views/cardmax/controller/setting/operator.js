layui.define(['admin', 'upload', 'table', 'form'], function (exports) {
    var $ = layui.$,
        table = layui.table,
        admin = layui.admin,
        setter = layui.setter;
    layui.admin.operator = {
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
        formSubmit: function (url, type, data, callback) {
            admin.req({
                url: setter.baseUrl + url
                , type: type
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000}, function () {
                        layer.closeAll();
                        callback()
                    });
                }
            });
        },
        confirm: function (msg, url, type, data, callback) {
            layer.confirm(msg, {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + url
                    , type: type
                    , data: data
                    , done: function (res) {
                        layer.msg(res.message, {icon: 1, time: 2000});
                        callback()
                    }
                });
            });

        },
        //获取表格中选中成员id
        getCheckList: function (tableName) {
            var idList = [];
            var checkStatus = table.checkStatus(tableName), data = checkStatus.data;
            for (var i in data) {
                idList.push(data[i].id);
            }
            return idList;
        },
        cancel: function () {
            $(document).off('click', '#cancel');
            $(document).on('click', '#cancel', function () {
                layer.close(admin.popup.index);
            })
        },
    };

    exports('setting/operator', layui.admin.operator)
})