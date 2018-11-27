layui.define(['admin', 'upload', 'table', 'form'], function (exports) {
    var $ = layui.$,
        table = layui.table,
        admin = layui.admin,
        setter = layui.setter,
        upload = layui.upload,
        form = layui.form;
    layui.admin.operation = {
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
        uploadImg: function () {
            upload.render({
                elem: '.upload_img'
                // ,field: 'content'
                , size: 1024
                , url: setter.baseUrl + '/public/oss-upload'
                , before: function (obj) {
                    //预读本地文件示例，不支持ie8
                    // current_id = $(this)[0].item.data('id');
                }
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000})
                    // console.log($('.banner_img' + current_id))
                    $('.banner_img').attr('src', res.data.ossUrl); //图片链接（base64）
                    $('.content').val(res.data.ossUrl);
                }
            });
        },
        formSubmit: function (url,type, data) {
            admin.req({
                url: setter.baseUrl + url
                , type: type
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000},function(){
                        window.history.go(-1);
                    });
                }
            });
        },
        deleteList: function (url, data,callback) {
            layer.confirm('确认删除选中的记录么？', {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + url
                    , type: 'delete'
                    , data: data
                    , done: function (res) {
                        layer.msg(res.message, {icon: 1, time: 1000},function(){
                            callback()
                        });
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

    exports('operation/operation', layui.admin.operation)
})