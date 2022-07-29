<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script lang="ts">
export default {
  data() {
    return {
      url: process.env.VUE_APP_URL,
      version: process.env.VUE_APP_VERSION,
    }
  },
}
</script>
