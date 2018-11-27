/**

 @Name：全局配置
 @Author：贤心
 @Site：http://www.layui.com/admin/
 @License：LPPL（layui付费产品协议）
    
 */
 
layui.define(['laytpl', 'layer', 'element', 'util'], function(exports){
  exports('setter', {
    container: 'LAY_app' //容器ID
    ,base: layui.cache.base //记录layuiAdmin文件夹所在路径
    ,views: layui.cache.base + 'views/' //视图所在目录
    ,entry: 'index' //默认视图文件名
    ,engine: '.html' //视图文件后缀名
    ,pageTabs: true //是否开启页面选项卡功能
    
    ,name: 'layuiAdmin Pro'
    ,tableName: 'layuiAdmin' //本地存储表名
    ,MOD_NAME: 'admin' //模块事件名
    
    ,debug: true //是否开启调试模式。如开启，接口异常时会抛出异常 URL 等信息
    
    ,interceptor: true //是否开启未登入拦截
    
    //自定义请求字段
    ,request: {
      tokenName: 'Authorization' //自动携带 token 的字段名。可设置 false 不携带。
    }
    
    //自定义响应字段
    ,response: {
      // statusName: 'status' //数据状态的字段名称
      // ,statusCode: {
      //   ok: 'S' //数据状态一切正常的状态码
      //   ,error: 'F' //数据状态一切正常的状态码
      //   ,logout: 1001 //登录状态失效的状态码
      // }
      msgName: 'massage' //状态信息的字段名称
      ,dataName: 'data' //数据详情的字段名称
    }

    //自定义表格响应字段
    ,responseTable: {
        statusName: '' //数据状态的字段名称，默认：code
        ,statusCode: null //成功的状态码，默认：0
        ,msgName: '' //状态信息的字段名称，默认：msg
        ,countName: 'count' //数据总数的字段名称，默认：count
        ,dataName: 'data'
    }
    ,pageTable: {
          layout: ['prev', 'page', 'next', 'count', 'skip'] //自定义分页布局
          ,groups: 5 //只显示 5 个连续页码
    }
    ,limit: 15

    ,baseUrl: '/v1/card-max/admin'
    
    //独立页面路由，可随意添加（无需写参数）
    ,indPage: [
      '/auth/login' //登入页
      ,'/auth/forget' //找回密码
      ,'/template/tips/lock' //平台关闭
    ]
    
    //扩展的第三方模块
    ,extend: [
      'echarts', //echarts 核心包
      'echartsTheme', //echarts 主题
      'Sortable',//拖拽排序
      'zTree',//拖拽排序
      'distpicker',
      'ChineseDistricts'
    ]
    
    //主题配置
    ,theme: {
      //配色方案，如果用户未设置主题，第一个将作为默认
      color: [/*{
        main: '#20222A' //主题色
        ,selected: '#009688' //选中色
        ,alias: 'default' //默认别名
      },{
        main: '#03152A'
        ,selected: '#3B91FF'
        ,alias: 'dark-blue' //藏蓝
      },{
        main: '#2E241B'
        ,selected: '#A48566'
        ,alias: 'coffee' //咖啡
      },{
        main: '#50314F'
        ,selected: '#7A4D7B'
        ,alias: 'purple-red' //紫红
      },{
        main: '#344058'
        ,logo: '#1E9FFF'
        ,selected: '#1E9FFF'
        ,alias: 'ocean' //海洋
      },{
        main: '#3A3D49'
        ,logo: '#2F9688'
        ,selected: '#5FB878'
        ,alias: 'green' //墨绿
      },{
        main: '#20222A'
        ,logo: '#F78400'
        ,selected: '#F78400'
        ,alias: 'red' //橙色
      },{
        main: '#28333E'
        ,logo: '#AA3130'
        ,selected: '#AA3130'
        ,alias: 'fashion-red' //时尚红
      },{
        main: '#24262F'
        ,logo: '#3A3D49'
        ,selected: '#009688'
        ,alias: 'classic-black' //经典黑
      },*/{
        main: '#22222a'
        ,logo: '#22222a'
        ,selected: '#0fcfbb'
        ,alias: 'classic-black' //经典黑
      }]
    }
  });
});
