layui.define(['admin', 'upload', 'table', 'form'], function (exports) {
    var $ = layui.$,
        table = layui.table,
        admin = layui.admin,
        setter = layui.setter,
        form = layui.form;
    layui.admin.message = {
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
        repairExport: function () {
            form.on('submit(LAY-exportRepair)', function (obj) {
                var time = obj.field.time.split('至')
                if(time.length > 1){
                    obj.field.beginTime = time[0];
                    obj.field.endTime = time[1];
                }else{
                    obj.field.beginTime = '';
                    obj.field.endTime = '';
                }
                delete obj.field.time;
                var href = setter.baseUrl + '/system-set/repairs-export?Authorization='
                    + layui.data('layuiAdmin').Authorization + '&operatorId=' + obj.field.operatorId + '&deviceNo=' + obj.field.deviceNo + '&beginTime=' + obj.field.beginTime + '&endTime=' + obj.field.endTime;
                window.location.href = href;
            });
        },
        feedbackExport: function () {
            form.on('submit(LAY-exportFeedback)', function (obj) {
                var time = obj.field.time.split('至');
                if(time.length > 1){
                    obj.field.beginTime = time[0];
                    obj.field.endTime = time[1];
                }else{
                    obj.field.beginTime = '';
                    obj.field.endTime = '';
                }
                delete obj.field.time;
                var href = setter.baseUrl + '/system-set/news/feedback-export?Authorization='
                    + layui.data('layuiAdmin').Authorization + '&operatorId='
                    + obj.field.operatorId + '&beginTime=' + obj.field.beginTime + '&endTime=' + obj.field.endTime;
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

    exports('setting/message', layui.admin.message)
})