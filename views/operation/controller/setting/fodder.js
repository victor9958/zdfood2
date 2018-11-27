/**
 * Created by meng on 2018/6/28.
 */
layui.define(['admin', 'table','layer'], function (exports) {
    var $ = layui.$,
        admin = layui.admin,
        table = layui.table,
        setter = layui.setter,
        view = layui.view
    $(function(){
        getList();
        function getList(other) {    //other:其他搜索字段
            if (other == undefined) {
                other = '';
            }
            table.render({
                id: 'fodder_list'
                , elem: '.wechat #fodder_list'
                , url: setter.baseUrl + '/banner/index'
                , page: setter.pageTable
                , cellMinWidth: 80
                , loading: false
                , limit: setter.limit
                , response:setter.responseTable
                , cols: [[
                    {type: 'numbers', title: '序号', width: 40}
                    , {field: 'pv', title: 'PV', width: 120}
                    , {field: 'uv', title: 'UV', width: 120}
                    , {field: 'updated_at', title: '录入时间',width: 200}
                    , {field: 'url', title: '链接地址',width: 200}
                    , {field: 'banner_title', width: 150, title: '标题'}
                    , {title: '操作', width: 150, toolbar: '#fodder_option'}
                ]]
                , where: {
                    'bannerTitle': other
                }
                , skin: 'line'
            });
        }
        //查询
        $(document).off('click','#search');
        $(document).on('click','#search',function(){
            var banner_title=$('.layui-form input[name=banner_title]').val();
            getList(banner_title)
        });
        //新增
        $(document).off('click','#add');
        $(document).on('click','#add',function(){
            admin.popupCenter({
                id: 'LAY_fodderAdd'
                , title: '新增'
                , area: ['460px', '500px']
                , success: function () {
                    view(this.id).render('setting/fodder/fodder/addmodal');
                }
            });
        });
        //修改
        table.on('tool(fodder_list)', function(obj){
            var data = obj.data;
            layui.data('fodder', {
                key: 'fodder'
                ,value: {'editId':data.id}
            });
            if(obj.event=='edit'){
                admin.popupCenter({
                    id: 'LAY_fodderEdit'
                    , title: '修改'
                    , area: ['440px', '460px']
                    , success: function () {
                        view(this.id).render('setting/fodder/fodder/editmodal');
                    }
                });
            }
            if(obj.event=='banner'){
                location.href="/#/setting/fodder/wechat_detail"
            }

        });
        //banner管理
        /*table.on('tool(fodder_list)', function(obj){
            var data = obj.data;
            layui.data('fodder', {
                key: 'fodder'
                ,value: {'editId':data.id}
            });
            admin.req({
                url: setter.baseUrl + '/banner/add-model'
                , type: 'POST'
                ,data:obj.field
                , done: function (res) {
                    layer.close(admin.popup.index);
                    layer.msg(res.message, {icon: 1, time: 2000});
                    setTimeout(function(){
                        window.parent.location.reload();
                    },2000);
                }
            });
        });*/
        //导入
        $(document).off('click','#import');
        $(document).on('click','#import',function(){
            admin.popupCenter({
                id: 'LAY_fodderImport'
                , title: '导入'
                , area: ['580px', '430px']
                , success: function () {
                    view(this.id).render('setting/fodder/fodder/importmember');
                }
            });
        });
    })
    exports('setting/fodder', {})
})