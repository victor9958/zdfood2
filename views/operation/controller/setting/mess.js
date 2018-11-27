/**
 * Created by meng on 2018/6/29.
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
            var other=other;
            table.render({
                id: 'mess_list'
                , elem: '.wechat #mess_list'
                , url: setter.baseUrl + '/templet/index'
                , page: setter.pageTable
                , cellMinWidth: 80
                , loading: false
                , limit: setter.limit
                , response:setter.responseTable
                , cols: [[
                    {type: 'numbers', title: '序号', width: 60}
                    , {field: 'id', title: '主键', width: 60}
                    , {field: 'wechat_name', title: '公众号名称', width: 120}
                    , {field: 'model_name', title: '应用模块', width: 120}
                    , {field: 'model_position', title: '模块位置',width: 120}
                    , {field: 'app_id', title: 'appId',width: 150}
                    , {field: 'movement_num', title: '推送量',width: 100}
                    , {field: 'click_num', width: 100, title: '点击率'}
                    , {field: 'templet_id', width: 200, title: '模板ID'}
                    , {field: 'created_at', width: 200, title: '生效日期'}
                    , {field: 'jump_url', width: 200, title: '跳转链接'}
                    , {field: 'status',title: '操作', width: 200, toolbar: '#mess_option'}
                ]]
                , where: {
                    'templetTitle': other
                }
                , skin: 'line'
            });
        }
        //查询
        $(document).off('click','#search');
        $(document).on('click','#search',function(){
            var data={};
            var templet_id=$('.layui-form #templet_id').val();
            var jump_url=$('.layui-form #jump_url').val();
            var wechat_name=$('.layui-form #wechat_name').val();
            var model_name=$('.layui-form #model_name').val();
            var model_position=$('.layui-form #model_position').val();
            data.templet_id=templet_id;
            data.jump_url=jump_url;
            data.wechat_name=wechat_name;
            data.model_name=model_name;
            data.model_position=model_position;
            getList(data);
        });
        //新增
        $(document).off('click','#add');
        $(document).on('click','#add',function(){
            admin.popupCenter({
                id: 'LAY_fodderAdd'
                , title: '新增'
                , area: ['440px', '460px']
                , success: function () {
                    view(this.id).render('setting/mess/mess/addmodal');
                }
            });
        });
        //导入
        $(document).off('click','#import');
        $(document).on('click','#import',function(){
            admin.popupCenter({
                id: 'LAY_fodderImport'
                , title: '导入'
                , area: ['580px', '430px']
                , success: function () {
                    view(this.id).render('setting/mess/mess/importmember');
                }
            });
        });
        table.on('tool(mess_list)', function(obj){
            var data = obj.data;
            layui.data('mess', {
                key: 'mess'
                ,value: {'templetId':data.id}
            });
            if(obj.event=='mess_edit'){
                admin.popupCenter({
                    id: 'LAY_messEdit'
                    , title: '修改'
                    , area: ['440px', '460px']
                    , success: function () {
                        view(this.id).render('setting/mess/mess/editmodal');
                    }
                });
            }
            if(obj.event=='mess_block'){
                if(data.status=='0'){
                    var block_title='停用';
                    var block_html='<div style="text-align: center;line-height:88px;">是否停用当前消息模板</div>';
                    var block_status='1';
                }else {
                    var block_title='启用';
                    var block_html='<div style="text-align: center;line-height:88px;">是否启用·当前消息模板</div>';
                    var block_status='0';
                }
                layer.open({
                    title:block_title
                    ,content: block_html
                    ,btn:["确定","取消"]
                    ,area: ['460px', '230px']
                    ,yes: function(index, layero){
                        admin.req({
                            url: setter.baseUrl + '/templet/changeStatus'
                            , type: 'POST'
                            , data: {"id":data.id,"status":block_status}
                            , done: function (res) {
                                layer.close(index);
                                layer.msg(res.message, {icon: 1, time: 2000});
                                getList();
                                //location.reload()
                            }
                        });
                    }
                });
            }
            if(obj.event=='mess_detail'){
                location.href="/#/setting/mess/wechat_detail"
            }

        });
    })
    exports('setting/mess', {})
})