layui.define(['admin', 'upload', 'table', 'form'], function (exports) {
    var $ = layui.$,
        table = layui.table,
        admin = layui.admin,
        setter = layui.setter,
        form = layui.form,
        upload = layui.upload;
    layui.admin.device = {
        tableData: function (tableName,option, query, deviceType) {
            layer.closeAll('tips');
            query.deviceType = deviceType;
            table.render({
                id: tableName
                , elem: '#' + tableName
                , url: setter.baseUrl + '/device-manage/devices' //模拟接口
                , page: setter.pageTable
                , cellMinWidth: 100
                , loading: false
                , limit: setter.limit
                , response: setter.responseTable
                , cols: [[
                    {type: 'checkbox', fixed: 'left'}
                    , {type: 'numbers', title: '序号', width: 40}
                    , {field: 'device_name', title: '设备编号'}
                    , {field: 'title', title: '设备类型', width: 120}
                    , {field: 'siteTitle', title: '设备位置'}
                    , {field: 'operatorName', title: '运营商', width: 120}
                    , {field: 'statusTitle', title: '设备状态', width: 100}
                    , {field: 'version', title: '版本号'}
                    , {title: '操作', width: 200, toolbar: option}
                ]]
                , where: query
                , skin: 'line'
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
        deviceSwith: function (url, data, msg, callback) {
            layer.confirm(msg, {icon: 3, title: '提示'}, function (index) {
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + url
                    , type: 'put'
                    , data: data
                    , done: function (res) {
                        layer.msg(res.message, {icon: 1, time: 1000},function(){
                            callback()
                        });
                    }
                });
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
                    current_id = $(this)[0].item.data('id');
                }
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000})
                    console.log($('.img_' + current_id))
                    $('.img_' + current_id).attr('src', res.data.ossUrl).show().siblings().hide(); //图片链接（base64）
                    $('.content_' + current_id).val(res.data.ossUrl);
                }
            });
        },
        deviceSetting: function (url, data) {
            admin.req({
                url: setter.baseUrl + url
                , type: 'put'
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 2000});
                }
            });
        },
        deviceApply: function (data, callback) {
            admin.req({
                url: setter.baseUrl + '/device-manage/devices-apply'
                , type: 'post'
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000},function(){
                        layer.closeAll();
                        callback();
                    });
                }
            });
        },
        deviceCheck: function (data, callback) {
            admin.req({
                url: setter.baseUrl + '/device-manage/check-devices-active'
                , type: 'put'
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000},function(){
                        callback()
                    });
                }
            });
        },
        deviceAlive: function (devceId, callback) {
            admin.req({
                url: setter.baseUrl + '/device-manage/device/' + devceId + '/init-upgrade'
                , type: 'put'
                , data: {}
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000},function(){
                        callback()
                    });
                }
            });
        },
        deviceRefresh: function (data, callback) {
            admin.req({
                url: setter.baseUrl + '/device-manage/devices-upgrade'
                , type: 'put'
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000},function(){
                        callback()
                    });
                }
            });
        },
        deviceQuery: function (data, callback) {
            admin.req({
                url: setter.baseUrl + '/device-manage/devices-version'
                , type: 'get'
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000},function(){
                        callback()
                    });
                }
            });
        },
        setLocation: function (data, callback) {
            admin.req({
                url: setter.baseUrl + '/device-manage/devices/site'
                , type: 'put'
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000},function(){
                        layer.closeAll();
                        callback()
                    });
                }
            });
        },
        deviceExport: function (filter,url,deviceType) {
            form.on(filter, function (obj) {
                var href = setter.baseUrl + url +'?Authorization='
                    + layui.data('layuiAdmin').Authorization + '&deviceType=' + deviceType + '&operatorId='
                    + obj.field.operatorId + '&status=' + obj.field.status + '&deviceNo=' + obj.field.deviceNo;
                window.location.href = href;
            });
        },
        useLogExport: function (deviceType) {
            form.on('submit(LAY-exportUseLog)', function (obj) {
                var time = obj.field.time.split('至');
                if(time.length > 1){
                    obj.field.beginTime = time[0];
                    obj.field.endTime = time[1];
                }else{
                    obj.field.beginTime = '';
                    obj.field.endTime = '';
                }
                delete obj.field.time;
                var href = setter.baseUrl + '/device-manage/use-records-export?Authorization='
                    + layui.data('layuiAdmin').Authorization + '&deviceType=' + deviceType + '&operatorId='
                    + obj.field.operatorId + '&beginTime=' + obj.field.beginTime+ '&endTime=' + obj.field.endTime + '&deviceNo=' + obj.field.deviceNo;
                window.location.href = href;
            });
        },
        useLog: function (tableName, query, deviceType) {
            layer.closeAll('tips');
            query.deviceType = deviceType;
            table.render({
                id: tableName
                , elem: '#' + tableName
                , url: setter.baseUrl + '/device-manage/use-records' //模拟接口
                , page: setter.pageTable
                , cellMinWidth: 100
                , loading: false
                , limit: setter.limit
                , response: setter.responseTable
                , cols: [[
                    {type: 'numbers', title: '序号', width: 60}
                    , {field: 'device_name', title: '设备编号'}
                    , {field: 'device_title', title: '设备类型'}
                    , {field: 'operatorName', title: '运营商'}
                    , {field: 'preAmountFormat', title: '订单金额'}
                    , {field: 'amountFormat', title: '实付金额'}
                    , {field: 'payTypeTitle', title: '支付方式'}
                    , {field: 'created_at', title: '时间', width: 240}
                    , {title: '操作', width: 100, toolbar: '#waterLogOption'}
                ]]
                , where: query
                , skin: 'line'
            });
        },
        importDevice: function (elem,deviceType) {
            upload.render({
                elem: elem //绑定元素
                , headers: {'Authorization': layui.data(setter.tableName)[setter.request.tokenName]}
                , url: setter.baseUrl + '/device-manage/firmware-Upload' //上传接口
                , accept: 'file' //允许上传的文件类型
                , before: function (obj) {
                    load = layer.load(); //上传loading
                }
                , done: function (res) {
                    layer.close(load);
                    // layer.msg(res.message, {icon: 1, time: 2000});
                    admin.req({
                        url: setter.baseUrl + '/device-manage/firmware-http-url' //实际使用请改成服务端真实接口
                        , type: 'put'
                        , data: {'filePath': res.data.filePath,deviceType:deviceType}
                        , done: function (res) {
                            layer.msg(res.message, {icon:1, time: 2000});
                        }
                    });
                }
                , error: function () {
                    //请求异常回调
                    layer.close(admin.popup.index);
                }
            });
        },
        setMode: function (data, callback) {
            admin.req({
                url: setter.baseUrl + '/device-manage/devices/mode'
                , type: 'put'
                , data: data
                , done: function (res) {
                    layer.msg(res.message, {icon: 1, time: 1000},function(){
                        callback()
                    });
                }
            });
        },
        cancel: function () {
            $(document).off('click', '#cancel');
            $(document).on('click', '#cancel', function () {
                layer.close(admin.popup.index);
            })
        },
    };

    exports('device/device', layui.admin.device)
})