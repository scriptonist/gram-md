<template>
<div>
<div class="row" style="margin-top:1.0em">
    <div id="md" class="col">
        <textarea style="height:auto" rows="16" class='form-control card' v-model='md_text'></textarea>
    </div>
</div>
<div class="col-md-6 space-button">
     <button v-bind:class="saveButtonClasses" v-on:click="save">{{saveButtonText}}</button>
</div>

<div class="row space-preview">
    <div id="preview" class="col-sm">
        <div>
             <div class = "well pre-scrollable text-left card"  v-html='previewText' ></div>
        </div>
    </div>
</div>
</div>
</template>

<script>
let marked = require('marked')
export default {
  name: 'Edit',
  data () {
    return {
      md_text: '# Hello',
      saveButtonClasses: {
        'btn-danger': false,
        'btn-primary': true,
        'btn-success': false,
        'btn float-left': true
      },
      saveButtonText: 'Save'
    }
  },
  methods: {
    save () {
      fetch('/save', {
        body: this.md_text,
        method: 'POST',
        mode: 'cors'
      }).then((response) => {
        if (response.status === 200) {
          this.saveButtonText = 'Saved'
          this.saveButtonClasses['btn-success'] = true
          this.saveButtonClasses['btn-primary'] = false
          setInterval(() => {
            this.saveButtonText = 'save'
            this.saveButtonClasses['btn-success'] = false
            this.saveButtonClasses['btn-primary'] = true
          }, 2000)
        }
      }).catch(err => {
        if (err) {
          this.saveButtonText = 'save failed'
          this.saveButtonClasses['btn-danger'] = true
          this.saveButtonClasses['btn-primary'] = false
          setInterval(() => {
            this.saveButtonText = 'save'
            this.saveButtonClasses['btn-danger'] = false
            this.saveButtonClasses['btn-primary'] = true
          }, 2000)
        }
      })
    }
  },
  computed: {
    previewText () {
      marked.setOptions({
        renderer: new marked.Renderer(),
        gfm: true,
        tables: true,
        breaks: true,
        pedantic: false,
        sanitize: true,
        smartLists: true,
        smartypants: false
      })
      return marked(this.md_text)
    }
  }
}

</script>

<style>
.space-button{
  margin-top:0.50em
}
.space-preview{
  margin-top:5em
}
</style>
