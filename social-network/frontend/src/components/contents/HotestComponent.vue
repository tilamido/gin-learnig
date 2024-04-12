<template>
    <ul class="tab-container" id="tabContainer" ref="tabContainer">
      <li class="tab-item" v-for="(item, index) in pbList" :key="index">
        <div><img :src="item.url" /></div>
        <div>
          <p class="content">{{ item.content }}</p>
        </div>
      </li>
    </ul>
  </template>
  
  <script>
  var defaults = {
    width: 360,
    delay: 100,
    repeatShow: false
  };
  export default {
    data() {
      return {
        config: {} ,
        //瀑布流数据
        pbList: [
          {
            url: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg3.doubanio.com%2Fview%2Fphoto%2Fm%2Fpublic%2Fp2650049201.jpg&refer=http%3A%2F%2Fimg3.doubanio.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1664935370&t=d4bf3e4d352c277a1bdebfcc8fda959f",
            title: "标题",
            content:
              "描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分",
            money: "68.50",
            sales_volume: "281",
          },
          {
            url: "https://img1.baidu.com/it/u=2911909188,130959360&fm=253&fmt=auto&app=138&f=JPEG?w=440&h=641",
            title: "标题",
            content: "描述部分描述部分描述部分描述部分描述部分描述部分描述部分",
            money: "35.00",
            sales_volume: "1221",
          },
          {
            url: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg3.doubanio.com%2Fview%2Fphoto%2Fm%2Fpublic%2Fp2650049201.jpg&refer=http%3A%2F%2Fimg3.doubanio.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1664935370&t=d4bf3e4d352c277a1bdebfcc8fda959f",
            title: "标题",
            content:
              "描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分",
            money: "68.50",
            sales_volume: "281",
          },
          {
            url: "https://img1.baidu.com/it/u=2911909188,130959360&fm=253&fmt=auto&app=138&f=JPEG?w=440&h=641",
            title: "标题",
            content: "描述部分描述部分描述部分描述部分描述部分描述部分描述部分",
            money: "35.00",
            sales_volume: "1221",
          },
          {
            url: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fimg3.doubanio.com%2Fview%2Fphoto%2Fm%2Fpublic%2Fp2650049201.jpg&refer=http%3A%2F%2Fimg3.doubanio.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1664935370&t=d4bf3e4d352c277a1bdebfcc8fda959f",
            title: "标题",
            content:
              "描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分描述部分",
            money: "68.50",
            sales_volume: "281",
          },
        ],
        observer: null,
      };
    },
    mounted() {
      this.config = { ...defaults };
      // console.log(this.config);
      // this.waterFall("#tabContainer", ".tab-item"); //实现瀑布流
      // // 窗口变化自适应布局
      // window.onresize = () => {
      //   return (() => {
      //     this.waterFall("#tabContainer", ".tab-item");
      //   })();
      // };
      // 使用 nextTick 确保 DOM 元素已经渲染完毕
      this.$nextTick(() => {
        this.waterFall("#tabContainer", ".tab-item");
        // 使用 ResizeObserver 监听尺寸变化
        this.observer = new ResizeObserver(() => {
          this.waterFall("#tabContainer", ".tab-item");
        });
        this.observer.observe(this.$refs.tabContainer); // 监听 tabContainer
        console.log(this);
      });
    },
    methods: {
      /**
       * @param { string } wrapIdName    容器id(或class)名称
       * @param { string } contentIdName 容器中内容项id(或class)名称
       * @param { number } column        容器中内容展示列数 手机的话建议改为2
       * @param { number } columnGap     容器中 列 间隔距离 默认为20
       * @param { number } rowGap        容器中 行 间隔距离 默认为20
       */
  
      //瀑布流方法：通过拿到dom循环，给每一个dom添加对应的定位位置排列出瀑布流布局。
      //通过判断列的高度，来把下一个盒子放在最短的地方补上
      waterFall(
        wrapIdName,
        contentIdName,
        // columns = 5,
        columnGap = 20,
        rowGap = 20
      ) {
        const wrapElement = document.querySelector(wrapIdName);
        if (!wrapElement || wrapElement.offsetWidth <= 8) {
          return;
        }
        
        // 获得内容可用宽度（去除滚动条宽度）
        const wrapContentWidth =
          wrapElement.offsetWidth - 8;
        
        // console.log(wrapContentWidth);
          
        var columns = parseInt(wrapContentWidth / this.config.width);
        // 间隔空白区域
        const whiteArea = (columns - 1) * columnGap;
  
        // 得到每列宽度(也即每项内容宽度)
        const contentWidth = parseInt((wrapContentWidth - whiteArea) / columns);
  
        // 得到内容项集合
        const contentList = document.querySelectorAll(contentIdName);
  
        // 成行内容项高度集合
        const lineConentHeightList = [];
  
        for (let i = 0; i < contentList.length; i++) {
          // 动态设置内容项宽度
          contentList[i].style.width = contentWidth + "px";
  
          // 获取内容项高度
          const height = contentList[i].clientHeight;
  
          if (i < columns) {
            // 第一行按序布局
            contentList[i].style.top = 0;
            contentList[i].style.left = contentWidth * i + columnGap * i + "px";
  
            // 将行高push到数组
            lineConentHeightList.push(height);
          } else {
            // 其他行
            // 获取数组最小的高度 和 对应索引
            let minHeight = Math.min(...lineConentHeightList);
            let index = lineConentHeightList.findIndex(
              (listH) => listH === minHeight
            );
  
            contentList[i].style.top = minHeight + rowGap + "px";
            contentList[i].style.left = (contentWidth + columnGap) * index + "px";
  
            // 修改最小列的高度 最小列的高度 = 当前自己的高度 + 拼接过来的高度 + 行间距
            lineConentHeightList[index] += height + rowGap;
          }
        }
      },
      beforeDestroy() {
        // console.log("destroy");
        // 在组件销毁之前停止观察
        if (this.observer) {
          this.observer.disconnect();
        }
      },
    },
  };
  </script>
  
  <style scoped>
  * {
    margin: 0;
    padding: 0;
  }
  ul{
    list-style-type: none;
    padding-left: 0;
  }
  /* 最外层 */
  .tab-container {
    position: relative;
    margin: 0 auto;
  }
  /* 每个小盒子 */
  .tab-container .tab-item {
    position: absolute;
    height: auto;
    border: 1px solid #ccc;
    /* box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04); */
    box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1);
    background: white;
    /* 元素不能中断显示 */
    break-inside: avoid;
    text-align: center;
  }
  
  .tab-container .tab-item:hover {
    transform: translateY(-6px);
    box-shadow: 0 30px 50px rgba(0, 0, 0, 0.3);
    transition: all 0.3s
}
  .tab-container .tab-item img {
    width: 100%;
    height: auto;
  }
  /* 描述 */
  .content {
    line-height: 25px;
    text-align: left;
    color: #5c5c5c;
    font-size: 14px;
    /* margin-top: 10px; */
    padding: 10px 10px 0 10px;
    overflow: hidden;
    text-overflow: ellipsis;
    /* 将对象作为弹性伸缩盒子模型显示 */
    display: -webkit-box;
    /* 限制在一个块元素显示的文本的行数 */
    /* -webkit-line-clamp 其实是一个不规范属性，使用了WebKit的CSS扩展属性，该方法适用于WebKit浏览器及移动端；*/
    -webkit-line-clamp: 2;
    /* 设置或检索伸缩盒对象的子元素的排列方式 */
    -webkit-box-orient: vertical;
  }
  </style>
  