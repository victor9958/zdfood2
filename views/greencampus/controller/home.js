layui.define(function (exports) {
    //数据概览
    layui.use(['echarts', 'form', 'table', 'laydate'], function () {
        var $ = layui.$
            , echarts = layui.echarts
            , setter = layui.setter
            , admin = layui.admin
            , table = layui.table
            , laydate = layui.laydate
            , form = layui.form;

        var countType = '', operatorId = '';
        var query = {};
        var echartsApp = [], options = []
            , elemDataView = $('.echarts_elem')
            , renderDataView = function (index) {
            echartsApp[index] = echarts.init(elemDataView[index], layui.echartsTheme);
            echartsApp[index].setOption(options[index]);
        };


        //没找到DOM，终止执行
        if (!elemDataView[0]) return;
        // renderDataView(0);
        // renderDataView(1);
        // renderDataView(2);
        // renderDataView(3);

        $(window).resize(function () {
            $.each(echartsApp, function (i, k) {
                echartsApp[i].resize();
            })
        })

      

        function dataBar(data) {
            var line_xAxis, line_yAxis = [];
            admin.req({
                url: setter.baseUrl + '/graph/count/week-orders'
                , type: 'get'
                , data: {}
                , done: function (res) {
                    var unit;
                    // if (data.countType != undefined && data.countType === 'amount') {
                    //     unit = '元'
                    // } else {
                    //     unit = '笔'
                    // }
                    line_xAxis = res.data.xAxis;

                    if (res.data.yAxis.length === 0) {
                        $('.echarts_empty').eq(0).html('该时间暂无数据').show().siblings().hide();
                        // $('.count').val(res.data.totalCount + unit)
                    } else {
                        $.each(res.data.yAxis, function (index, item) {
                            console.log(111);
                            console.log(item);
                            item.type = 'bar';
                            line_yAxis.push(item)
                            console.log(line_yAxis);
                        })
                        options[0] = {
                            // color: ['#67E0E3'],
                            tooltip: {
                                trigger: 'axis',
                                axisPointer: {            // 坐标轴指示器，坐标轴触发有效
                                    type: 'shadow'        // 默认为直线，可选为：'line' | 'shadow'
                                }
                            },
                            legend: {
                                data:['支出','收入']
                            },
                            grid: {
                                x: 60,
                                y: 10,
                                x2: 10,
                                y2: 30

                            },
                            xAxis: [
                                {
                                    type: 'category',
                                    axisLabel: {
                                        interval: 0,
                                        // rotate: 40
                                    },
                                    data: line_xAxis
                                }
                            ],
                            yAxis: [
                                {
                                    type: 'value'
                                }
                            ],
                            series: line_yAxis
                        }

                        renderDataView(0);
                        $('.echarts_empty').eq(0).hide().siblings().show();
                        $('.count').val(res.data.totalCount + unit)
                    }

                }
            })
        }

        // function dataPie(data) {
        //     if(countType){
        //         delete data.countType
        //     }
        //     admin.req({
        //         url: setter.baseUrl + '/graph/count/week-orders'
        //         , type: 'get'
        //         , data: data
        //         , done: function (res) {
        //             if (res.data.length === 0) {
        //                 $('.echarts_empty').eq(1).html('该时间暂无数据').show().siblings().hide();
        //             } else {
        //                 options[1] = {
        //                     tooltip: {
        //                         trigger: 'item',
        //                         formatter: "{a} <br/>{b}: {c}笔 ({d}%)"
        //                     },
        //                     legend: {
        //                         show: false,
        //                         data: []
        //                     },
        //                     series: [
        //                         {
        //                             name: '支付方式',
        //                             type: 'pie',
        //                             radius: ['55%', '80%'],
        //                             avoidLabelOverlap: false,
        //                             label: {
        //                                 normal: {
        //                                     show: false,
        //                                     position: 'center',
        //                                     formatter: function (param) {
        //                                         var str = param.name + '\n' + param.value + '笔';
        //                                         return str
        //                                     }
        //                                 },
        //                                 emphasis: {
        //                                     show: true,
        //                                     textStyle: {
        //                                         fontSize: '20',
        //                                     }
        //                                 }
        //                             },
        //                             labelLine: {
        //                                 normal: {
        //                                     show: false
        //                                 }
        //                             },
        //                             data: res.data

        //                         }
        //                     ]
        //                 }
        //                 renderDataView(1);
        //                 $('.echarts_empty').eq(1).hide().siblings().show();
        //             }

        //         }
        //     })
        // }

        dataBar();
        // dataPie();


        //监听侧边伸缩
        layui.admin.on('side', function () {
            setTimeout(function () {
                $.each(echartsApp, function (i, k) {
                    echartsApp[i].resize();
                })
            }, 300);
        });

    });


    exports('home', {})
});