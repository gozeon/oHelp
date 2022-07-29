<template>
  <div>
    <tinymce :id="id" :other_options="tinyOptions" v-model="editorContent" @editorChange="handleChange"></tinymce>
  </div>
</template>
<script>
import cloneDeep from 'lodash/cloneDeep'
import tinymce from 'vue-tinymce-editor'
export default {
  props: {
    id: {
      type: Number | String,
      required: true,
    },
    value: {
      type: Number | String,
      required: true,
    },
    height: {
      type: Number,
      default: 500,
    },
  },
  components: { tinymce },
  data() {
    return {
      editorContent: '',
      tinyOptions: {
        branding: false,
        height: this.height,
        paste_data_images: true,
        language_url: 'tinymce/zh_CN.js',
        media_live_embeds: true,
        images_upload_handler: (blobInfo, succFun, failFun) => {
          let params = new FormData()
          params.append('file', blobInfo.blob())

          this.axios
            .post('/upload', params, {
              headers: {
                'Content-Type': 'multipart/form-data',
              },
            })
            .then((res) => {
              const { status, msg, data } = res.data
              if (status === 0) {
                succFun(data.url)
              } else {
                failFun(msg)
              }
            })
        },
        file_picker_types: 'file image media',
        file_picker_callback: (callback, value, meta) => {
          //文件分类
          var filetype = '.pdf, .txt, .zip, .rar, .7z, .doc, .docx, .xls, .xlsx, .ppt, .pptx, .mp3, .mp4'
          //为不同插件指定文件类型及后端地址
          switch (meta.filetype) {
            case 'image':
              filetype = '.jpg, .jpeg, .png, .gif'
              break
            case 'media':
              filetype = '.mp3, .mp4'
              break
            case 'file':
            default:
          }

          //模拟出一个input用于添加本地文件
          var input = document.createElement('input')
          input.setAttribute('type', 'file')
          input.setAttribute('accept', filetype)
          input.click()
          input.onchange = (e) => {
            let params = new FormData()
            params.append('file', e.target.files[0])

            this.axios
              .post('/upload', params, {
                headers: {
                  'Content-Type': 'multipart/form-data',
                },
              })
              .then((res) => {
                const { status, msg, data } = res.data
                if (status === 0) {
                  if (meta.filetype == 'file') {
                    callback(data.url, { text: 'My text' })
                  }

                  if (meta.filetype == 'image') {
                    callback(data.url, { alt: 'My alt text' })
                  }

                  if (meta.filetype == 'media') {
                    callback(data.url, { source1: data.url, source2: data.url, poster: '' })
                  }
                } else {
                  callback('')
                }
              })
          }
        },

        templates: [
          { title: 'Some title 1', description: 'Some desc 1', content: 'My content' },
          { title: 'Some title 2', description: 'Some desc 2', url: 'tinymce/tpl/tpl.html' },
        ],
      },
    }
  },
  created() {
    this.editorContent = cloneDeep(this.value)
  },
  methods: {
    handleChange() {
      this.$emit('input', this.editorContent)
    },
  },
}
</script>
