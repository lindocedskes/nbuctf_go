// export const downloadImage = (imgsrc, name) => {
//   // 下载图片地址和图片名
//   console.log('imgsrc', imgsrc)
//   var image = new Image()
//   image.setAttribute('crossOrigin', 'anonymous')
//   image.onload = function () {
//     var canvas = document.createElement('canvas')
//     canvas.width = image.width
//     canvas.height = image.height
//     var context = canvas.getContext('2d')
//     context.drawImage(image, 0, 0, image.width, image.height)
//     var url = canvas.toDataURL('image/png') // 得到图片的base64编码数据

//     var a = document.createElement('a') // 生成一个a元素
//     var event = new MouseEvent('click') // 创建一个单击事件
//     a.download = name || 'photo' // 设置图片名称
//     a.href = url // 将生成的URL设置为a.href属性
//     a.dispatchEvent(event) // 触发a的单击事件
//   }
//   image.src = imgsrc
// }

//下载文件
export const downloadImage = (fileUrl, fileName) => {
  fetch(fileUrl)
    .then((resp) => resp.blob())
    .then((blob) => {
      const url = window.URL.createObjectURL(blob)
      const a = document.createElement('a')
      a.style.display = 'none'
      a.href = url
      a.download = fileName
      document.body.appendChild(a)
      a.click()
      window.URL.revokeObjectURL(url)
      document.body.removeChild(a)
    })
    .catch(() => alert('文件下载失败'))
}
