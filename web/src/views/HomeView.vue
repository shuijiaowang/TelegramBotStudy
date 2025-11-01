
<template>
 <div>
   <button @click.stop="getTextHandler">点击测试</button>
   <span>{{textResponse}}</span>
 </div>

</template>
<script setup>
  import {ref} from "vue";
  import {text} from "@/api/example.js";

  const textResponse = ref(null)
  textResponse.value=111

  const getTextHandler=async () => {
    const res = await text({})
    textResponse.value=res.data
  }
</script>