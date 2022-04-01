<template>
  <div class="slid-nav">
    <div ref="bgBlock" class="bg-block" />
    <div ref="itemList" class="slid-nav-item-list">
      <div
        v-for="(item, index) in list"
        :key="item.path"
        class="slid-nav-item"
        @mouseenter="(e) => mouseEnter(e, index)"
        @mouseleave="(e) => mouseLeave(e, index)"
        @click="click(item, index)"
      >
        {{ item.label }}
      </div>
    </div>
  </div>
</template>

<script>
/**
 * list => [{label, path}]
 * value => path
 */
export default {
  name: 'SlidNav',
  props: {
    list: {
      type: Array,
      default: () => [],
    },
    value: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      bgBlock: null,
      itemList: null,
      originIndex: 0,
      timer: null,
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.bgBlock = this.$refs.bgBlock
      this.itemList = this.$refs.itemList
      this.list.find((item, index) => {
        if (item.path === this.value) {
          this.originIndex = index
          return true
        }
        return false
      })
      this.rollback()
    },
    mouseEnter(e, index) {
      clearTimeout(this.timer)
      this.bgBlock.style.left = e.target.offsetLeft + 'px'
      this.itemList.children[this.originIndex].classList.remove(
        'slid-nav-item-selected'
      )
    },
    mouseLeave(e, index) {
      this.timer = setTimeout(() => {
        this.rollback()
      }, 100)
    },
    rollback() {
      this.bgBlock.style.left =
        this.itemList.children[this.originIndex].offsetLeft + 'px'
      this.itemList.children[this.originIndex].classList.add(
        'slid-nav-item-selected'
      )
    },
    click(item, index) {
      this.originIndex = index
      clearTimeout(this.timer)
      this.itemList.children[this.originIndex].classList.add(
        'slid-nav-item-selected'
      )
      this.$emit('update:value', item.path)
    },
  },
}
</script>

<style lang="less" scoped>
.slid-nav {
  position: relative;
  border-radius: 50px;
  background-color: #dfdfdf;
  padding: 0 2px;
  .slid-nav-item-list {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 5px;
    .slid-nav-item {
      z-index: 1;
      border-radius: 50px;
      padding: 5px 5px;
      user-select: none;
      cursor: pointer;
      text-align: center;
      min-width: 4em;
      box-sizing: content-box;
      color: #888;
      &:hover {
        color: white;
      }
    }
    .slid-nav-item-selected {
      color: white;
    }
  }
  .bg-block {
    position: absolute;
    left: 0;
    top: 50%;
    transform: translateY(-50%);
    width: 4em;
    height: 1em;
    box-sizing: content-box;
    padding: 5px 5px;
    background-color: #f63e46;
    border-radius: 50px;
    border: 2px solid #dfdfdf;
    transition: all 0.3s;
  }
}
</style>
