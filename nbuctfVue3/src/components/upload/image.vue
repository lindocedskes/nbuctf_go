<template>
  <div>
    <el-upload
      :action="`${path}/file/upload`"
      :show-file-list="false"
      :on-success="handleImageSuccess"
      :before-upload="beforeImageUpload"
      :multiple="false"
    >
      <el-button type="primary">图片压缩上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
import ImageCompress from '@/utils/image'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { fileUpload } from '@/api/fileUploadAndDownload.js'

defineOptions({
  name: 'UploadImage'
})

const emit = defineEmits(['on-success'])
//接受父组件传递的参数，使用defineProps
const props = defineProps({
  imageUrl: {
    type: String,
    default: ''
  },
  fileSize: {
    type: Number,
    default: 2048 // 2M 超出后执行压缩
  },
  maxWH: {
    type: Number,
    default: 1920 // 图片长宽上限
  }
})

const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()
const fullscreenLoading = ref(false)

const beforeImageUpload = async (file) => {
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  if (!isJPG && !isPng) {
    ElMessage.error('上传头像图片只能是 jpg或png 格式!')
    return false
  }

  const isRightSize = file.size / 1024 < props.fileSize
  if (!isRightSize) {
    // 压缩
    const compress = new ImageCompress(file, props.fileSize, props.maxWH)
    // return compress.compress()
    file = await compress.compress()
  }
  // 自定义上传，通过@/utils/request，向8889发送请求
  const formData = new FormData()
  formData.append('file', file)
  try {
    const response = await fileUpload(formData)
    if (response.code === 0) {
      uploadSuccess(response)
    } else {
      uploadError()
    }
  } catch (error) {
    uploadError()
  }

  // 阻止组件默认的上传行为
  return false
}

const handleImageSuccess = (res) => {
  const { data } = res
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

const uploadSuccess = (res) => {
  const { data } = res //解构赋值data=res.data
  if (data.file) {
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
    emit('on-success', data.file.url) //当文件上传成功，uploadSuccess 方法会被调用，然后它会触发 on-success 事件，并传递 data.file.url 作为参数。然后，在父组件中，uploadSuccess 方法会被执行，它的参数就是 data.file.url。
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}
</script>

<style lang="scss" scoped>
.image-uploader {
  border: 1px dashed #d9d9d9;
  width: 180px;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.image-uploader {
  border-color: #409eff;
}
.image-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.image {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
