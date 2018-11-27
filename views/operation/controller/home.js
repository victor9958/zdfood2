layui.define(function (exports) {
    //数据概览
    layui.use(['echarts', 'form'], function () {
        var $ = layui.$
            , echarts = layui.echarts
            , form = layui.form;


        var echartsApp = [], options = [
            //注册用户图像
            {
                title: {
                    show: false
                },
                tooltip: {
                    trigger: 'item',
                    formatter: "{a} <br/>{b}: {c}人 ({d}%)"
                },
                legend: {
                    show: false,
                    data: []
                },
                series: [{
                    name: '注册用户画像',
                    type: 'pie',
                    radius: '60%',
                    center: ['50%', '50%'],
                    data: [
                        {value: 3052, name: '外来人员'},
                        {value: 1610, name: '来宾'},
                        {value: 3200, name: '职工'},
                        {value: 535, name: '管理层'},
                        {value: 1700, name: '保洁'},
                        {value: 1500, name: '参观团'},
                        {value: 2700, name: '嘉宾'},
                        {value: 100, name: '领导'}
                    ]
                }]
            },
            //各类支付方式对比
            {
                tooltip: {
                    trigger: 'item',
                    formatter: "{a} <br/>{b}: {c}笔 ({d}%)"
                },
                legend: {
                    show: false,
                    data: []
                },
                series: [
                    {
                        name: '支付方式',
                        type: 'pie',
                        radius: ['40%', '60%'],
                        avoidLabelOverlap: false,
                        label: {
                            normal: {
                                show: false,
                                position: 'center',
                                formatter: function (param) {
                                    var str = param.name + '\n' + param.value + '笔';
                                    return str
                                }
                            },
                            emphasis: {
                                show: true,
                                textStyle: {
                                    fontSize: '20',
                                }
                            }
                        },
                        labelLine: {
                            normal: {
                                show: false
                            }
                        },
                        data: [
                            {value: 3052, name: '外来人员'},
                            {value: 1610, name: '来宾'},
                            {value: 3200, name: '职工'},
                            {value: 535, name: '管理层'},
                            {value: 1700, name: '保洁'}
                        ]
                    }
                ]
            },

            //平台整体订单变化趋势图
            // {
            //     title: {
            //         show: false
            //     },
            //     grid:{
            //         y:30,
            //         x2:5,
            //         y2:20,
            //     },
            //     tooltip : {
            //         trigger: 'item',
            //         formatter: "{b} <br/>{a}: {c}"
            //     },
            //     xAxis: {
            //         type: 'category',
            //         data: ['02/01', '02/02', '02/03', '02/04', '02/05', '02/06', '02/07']
            //     },
            //     yAxis: {
            //         type: 'value'
            //     },
            //     series: [{
            //         name: '分类名称',
            //         data: [820, 932, 901, 934, 1290, 1330, 1320],
            //         type: 'line',
            //         smooth: true
            //     },{
            //         name: '分类名称1',
            //         data: [720, 1032, 801, 834, 1390, 1130, 1420],
            //         type: 'line',
            //         smooth: true
            //     },{
            //         name: '分类名称2',
            //         data: [920, 932, 1001, 734, 1290, 1330, 1520],
            //         type: 'line',
            //         smooth: true
            //     }]
            // },
            {
                color: ['#FF9F7F','#8378EA'],
                tooltip: {
                    trigger: 'none',
                    axisPointer: {
                        type: 'cross'
                    }
                },
                legend: {
                    data: ['2015 降水量', '2016 降水量']
                },
                grid: {
                    y: 30,
                    x2: 5,
                    y2: 20,
                },
                xAxis: [
                    {
                        type: 'category',
                        axisTick: {
                            alignWithLabel: true
                        },
                        axisLine: {
                            onZero: false,
                            lineStyle: {
                                color: '#FF9F7F'
                            }
                        },
                        axisPointer: {
                            label: {
                                formatter: function (params) {
                                    return '降水量  ' + params.value
                                        + (params.seriesData.length ? '：' + params.seriesData[0].data : '');
                                }
                            }
                        },
                        data: ["2016-1", "2016-2", "2016-3", "2016-4", "2016-5", "2016-6", "2016-7", "2016-8", "2016-9", "2016-10", "2016-11", "2016-12"]
                    },
                    {
                        axisLine: {
                            onZero: false,
                            lineStyle: {
                                color: '#8378EA'
                            }
                        },
                        axisPointer: {
                            label: {
                                formatter: function (params) {
                                    return '降水量  ' + params.value
                                        + (params.seriesData.length ? '：' + params.seriesData[0].data : '');
                                }
                            }
                        },
                        show: false,
                        data: ["2015-1", "2015-2", "2015-3", "2015-4", "2015-5", "2015-6", "2015-7", "2015-8", "2015-9", "2015-10", "2015-11", "2015-12"]
                    }
                ],
                yAxis: [
                    {
                        type: 'value'
                    }
                ],
                series: [
                    {
                        name: '2015 降水量',
                        type: 'line',
                        xAxisIndex: 1,
                        smooth: true,
                        data: [2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3]
                    },
                    {
                        name: '2016 降水量',
                        type: 'line',
                        smooth: true,
                        data: [3.9, 5.9, 11.1, 18.7, 48.3, 69.2, 231.6, 46.6, 55.4, 18.4, 10.3, 0.7]
                    }
                ]
            },

            //各模块使用情况
            {
                color: ['#67E0E3'],
                tooltip: {
                    trigger: 'axis',
                    axisPointer: {            // 坐标轴指示器，坐标轴触发有效
                        type: 'shadow'        // 默认为直线，可选为：'line' | 'shadow'
                    }
                },
                grid: {
                    y: 30,
                    x2: 5,
                },
                xAxis: [
                    {
                        type: 'category',
                        axisLabel: {
                            interval: 0,
                            rotate: 40
                        },
                        data: ['洗衣支付', '充电支付', '宿舍订水', '班车订票', '电费缴费', '洗衣支付', '充电支付', '洗衣支付', '充电支付', '宿舍订水', '班车订票', '电费缴费', '洗衣支付', '充电支付'],
                    }
                ],
                yAxis: [
                    {
                        type: 'value'
                    }
                ],
                series: [
                    {
                        name: '直接访问',
                        type: 'bar',
                        barWidth: '60%',
                        data: [10, 52, 200, 334, 390, 330, 220, 10, 52, 200, 334, 390, 330, 220]
                    }
                ]
            },
        ]
            , elemDataView = $('.echarts_elem')
            , renderDataView = function (index) {
            echartsApp[index] = echarts.init(elemDataView[index], layui.echartsTheme);
            echartsApp[index].setOption(options[index]);
        };


        //没找到DOM，终止执行
        if (!elemDataView[0]) return;

        renderDataView(0);
        renderDataView(1);
        renderDataView(2);
        renderDataView(3);

        $(window).resize(function () {
            $.each(echartsApp, function (i, k) {
                echartsApp[i].resize();
            })
        })

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