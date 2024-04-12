<template>
  <div class="container">
    <div class="post-box">
      <div class="button-group">
        <input type="file" ref="fileInput" style="display: none" multiple @change="handleFileChange">
      </div>
      <textarea v-model="postContent" placeholder="分享新鲜事..."></textarea>
      <div v-if="imagePreviews.length > 0" class="image-previews">
        <div v-for="(preview, index) in imagePreviews" :key="index" class="image-preview">
          <img :src="preview.url" alt="图片预览">
          <button @click="removeImagePreview(index)" class="remove-button">×</button>
        </div>
      </div>
      <button @click="chooseImage">上传图片</button>
      <button @click="publishPost">发布</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      postContent: '', // 发布的内容
      imagePreviews: [], // 图片预览数组
      selectedFiles: [], // 选择的文件数组
    };
  },
  methods: {
    chooseImage() {
      this.$refs.fileInput.click();
    },
    handleFileChange(event) {
      const newFiles = Array.from(event.target.files);
      this.selectedFiles.push(...newFiles);
      newFiles.forEach(file => {
        const reader = new FileReader();
        reader.onload = () => {
          this.imagePreviews.push({
            url: reader.result,
            file: file
          });
        };
        reader.readAsDataURL(file);
      });
    },
    removeImagePreview(index) {
      this.imagePreviews.splice(index, 1);
      this.selectedFiles.splice(index, 1);
    },
    publishPost() {
      if (this.postContent.trim() !== '') {
        // 将发布的内容传递给后端
        const postData = {
          content: this.postContent,
          images: this.selectedFiles,
        };
        // 这里调用后端API将postData发送到后端进行处理
        console.log('发布的内容：', postData);
        // 清空发布框内容和图片预览
        this.postContent = '';
        this.imagePreviews = [];
        this.selectedFiles = [];
      }
    },
  },
};
</script>

<style>
.container {
  width: 600px;
  margin: 20px auto;
}

.post-box {
  margin-bottom: 20px;
  background-color: var(--sidebar-color);
  border: 1px solid var(--primary-color);
  padding: 20px;
  border-radius: 10px;
}

textarea {
  width: 100%; /* 占据整行宽度 */
  height: 100px;
  resize: none;
  margin-bottom: 10px;
  border: 1px solid var(--primary-color);
  padding: 10px;
  border-radius: 5px;
}

textarea:focus {
  border-color: var(--primary-color);
  outline: none;
  box-shadow: 0 0 6px rgba(255, 69, 0, 0.5);
}

.button-group {
  display: flex;
  margin-bottom: 10px;
}

button {
  padding: 10px 20px;
  margin-right: 10px;
  color: var(--text-color);
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: var(--tran-03);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

button:hover {
  color: #fff;
  background-color: var(--primary-color);
}

.image-previews {
  margin: 5px 0;
  display: flex;
  flex-wrap: wrap;
}

.image-preview {
  margin-right: 10px;
  margin-bottom: 10px;
  position: relative;
}

.image-preview img {
  max-width: 100%;
  max-height: 200px;
  margin-bottom: 5px;
}

.remove-button {
  position: absolute;
  top: 5px;
  right: 5px;
  background-color: rgba(0, 0, 0, 0.5);
  color: #fff;
  border: none;
  padding: 0;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  font-size: 12px;
  cursor: pointer;
}
</style>
