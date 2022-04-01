<template>
  <a-card :bordered="false">
       <Home />
  </a-card>
</template>
<script>

import Home from "./components/web/home"
import News from "./components/web/news"
export default {
    components:{
        Home,
        News
    },
    data() {
        return {

        }
    },
    methods: {
        callback(key) {
        console.log(key);
        },
    },
};
</script>
