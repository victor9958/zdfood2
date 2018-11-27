layui.define(['table','form','laydate'], function(exports){
    var $ = layui.$,
        view = layui.view,
        table = layui.table,
        form = layui.form,
        admin = layui.admin,
        laydate = layui.laydate,
        setter = layui.setter;

    $(function(){
        tableData();

        var start = laydate.render({
            elem: '#beginTime',
            done: function(value, date, endDate){
                end.config.min = {
                    year: date.year,
                    month: date.month - 1,
                    date: date.date,
                }; //开始日选好后，重置结束日的最小日期
                end.config.max = {
                    year: date.year,
                    month: date.month,
                    date: date.date,
                }; //开始日选好后，重置结束日的最小日期
            }
        });

        //结束时间
        var end = laydate.render({
            elem: '#endTime',
            done: function(value, date, endDate){
                start.config.min = {
                    year: date.year,
                    month: date.month - 2,
                    date: date.date,
                }; //结束日选好后，重置开始日的最大日期
                start.config.max = {
                    year: date.year,
                    month: date.month - 1,
                    date: date.date,
                }; //结束日选好后，重置开始日的最大日期
            }
        });


        form.on('submit(LAY-search-feedback)', function(obj){
            tableData(obj.field);
        });


        form.on('submit(LAY-export-feedback)', function(obj){
            var download_href = setter.baseUrl + '/news-center/feedback/export?beginTime='+obj.field.beginTime+'&endTime='+obj.field.endTime+'&queryString='+obj.field.queryString;
            $(this).attr('href',download_href);
        });

        table.on('tool(feedback_)', function(obj){
            var data = obj.data;
            layui.data('temporary', {
                key: 'temporary'
                ,value: {'feedback_id':data.id}
            });
            admin.popupCenter({
                id: 'LAY_adminPopupFeedbackDetail'
                , title: '详情'
                , area: ['430px', '520px']
                , success: function () {
                    view(this.id).render('setting/message/feedback/detail');
                }
            });
        });

        form.on('submit(LAY-feedback)', function(obj){
            var reply = $('#reply').val();
            if(!reply){
                layer.alert('反馈内容不能为空', {icon: 5});
            }else{
                admin.req({
                    url: setter.baseUrl + '/news-center/feedback/'+ layui.data('temporary').temporary.feedback_id +'/reply'
                    ,type: 'put'
                    ,data: {'replyContent':reply}
                    ,done: function(res){
                        layer.msg(res.message, {icon: 1,time: 2000});
                        layer.close(admin.popup.index);//关闭面板
                    }
                });
            }
        });

        //删除反馈
        $(document).off('click', '#delete-feedback');
        $(document).on('click', '#delete-feedback', function () {
            layer.confirm('是否确认删除这条反馈？', {icon: 3, title: '提示'},function(index){
                layer.close(index);
                admin.req({
                    url: setter.baseUrl + '/news-center/feedback/'+ layui.data('temporary').temporary.feedback_id
                    ,type: 'delete'
                    ,data: {}
                    ,done: function(res){
                        layer.msg(res.message, {icon: 1,time: 2000});
                        tableData()
                        layer.close(admin.popup.index);//关闭面板
                    }
                });
            })

        })

        //查看图片
        $(document).off('click', '.feedback_img');
        $(document).on('click', '.feedback_img', function () {
            var src = $(this).find('img').attr('src');
            var width = $(this).find('img').width();
            var height = $(this).find('img').height();
            layer.open({
                type: 1,
                title: false,
                closeBtn: 0,
                shadeClose: true,
                area: 'auto', //宽高
                maxWidth: 450,
                content: '<img width="100%" height="100%" greencampus="' + src + '"/>'
            });

        })

        //关闭弹窗
        $(document).off('click', '#cancel');
        $(document).on('click', '#cancel', function () {
            layer.close(admin.popup.index);
        })

        function tableData(field){
            table.render({
                id: 'feedback_'
                ,elem: '#feedback_'
                ,url: setter.baseUrl + '/news-center/feedback' //模拟接口
                ,page: setter.pageTable
                ,cellMinWidth: 100
                ,loading: false
                ,limit: setter.limit
                ,response: setter.responseTable
                ,cols: [[
                    {type: 'numbers',title:'序号',width: 60}
                    ,{width: 16, templet: '#isVisit'}
                    ,{field: 'content', title: '反馈内容'}
                    ,{field: 'user_mobile', width: 140, title: '手机号'}
                    ,{field: 'user_name', width: 140, title: '姓名'}
                    ,{field: 'created_at', width: 200, title: '时间'}
                    ,{field: 'reply_content', width: 100, title: '回复状态', templet: '#isReply'}
                    ,{title: '操作', width: 100, toolbar: '#feedback_option'}
                ]]
                ,where: field
                ,skin: 'line'
            });
        }
    })

    //对外暴露的接口
    exports('setting/feedbackmsg', {});
});