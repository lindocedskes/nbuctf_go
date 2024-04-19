<template>
  <div>
    <el-upload
      :action="`${path}/file/upload`"
      :before-upload="checkFile"
      :show-file-list="false"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      class="upload-btn"
    >
      <el-button type="primary">上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { isVideoMime, isImageMime } from '@/utils/image'
import { fileUpload } from '@/api/fileUploadAndDownload.js'

defineOptions({
  name: 'UploadCommon'
})
const emit = defineEmits(['on-success'])
const path = ref(import.meta.env.VITE_BASE_API)

const fullscreenLoading = ref(false)

const checkFile = async (file) => {
  fullscreenLoading.value = true
  const isLt5M = file.size / 1024 / 1024 < 5 // 5MB
  const isLt50M = file.size / 1024 / 1024 < 50 // 50MB
  const isVideo = isVideoMime(file.type)
  const isImage = isImageMime(file.type)
  const isZip = file.type === 'application/zip'

  let pass = true
  if (!isVideo && !isImage && !isZip) {
    ElMessage.error(
      '上传文件只能是zip格式，上传图片只能是 jpg,png,svg,webp 格式, 上传视频只能是 mp4,webm 格式!'
    )
    fullscreenLoading.value = false
    pass = false
  }
  if (!isLt50M && isVideo) {
    ElMessage.error('上传视频大小不能超过 50MB')
    fullscreenLoading.value = false
    pass = false
  }
  if (!isLt5M && isImage) {
    ElMessage.error('未压缩的上传图片大小不能超过 5M，请使用压缩上传')
    fullscreenLoading.value = false
    pass = false
  }
  //自定义上传，通过@/utils/request，向8889发送请求
  if (pass) {
    const formData = new FormData()
    formData.append('file', file)
    try {
      const response = await fileUpload(formData)
      console.log(response)
      if (response.code === 0) {
        console.log(1)
        uploadSuccess(response)
      } else {
        uploadError()
      }
    } catch (error) {
      console.log(3)
      uploadError()
    }
  }

  // Prevent the default upload behavior
  return false
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
