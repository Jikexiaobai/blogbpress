<template>
    <div class="container">
        <div class="warper" :style="{ width: design.width + 'px' }">
            <WidgetOne />
            <WidgetTwo />
            <WidgetThree />
            <WidgetFive />
            <WidgetFour />
        </div>
    </div>
</template>

<script>
import api from "@/api/index"
import { mapState } from "vuex"
import WidgetOne from "@/components/widget/class/widgetOne"
import WidgetTwo from "@/components/widget/class/widgetTwo"
import WidgetThree from "@/components/widget/class/widgetThree"
import WidgetFour from "@/components/widget/class/widgetFour"
import WidgetFive from "@/components/widget/class/widgetFive"
export default {
    components:{
        WidgetOne,
        WidgetTwo,
        WidgetThree,
        WidgetFour,
        WidgetFive
    },
    computed:{
        ...mapState(["design"]),
    },
}
</script>

<style lang="less" scoped>
.container{
    margin: 50px 0;
    min-height: 550px;
    display: flex;
    justify-content: center;
}
</style>