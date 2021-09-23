<template>
  <div class="searchbar">
    <van-search
      :placeholder="placeholder"
      :modelValue="modelValue"
      v-on:input="updateSearchQuery($event.target.value)"
      background="none"
      shape="round"
    >
      <template v-slot:left-icon>
        <img id="search-icon" :src="SearchIcon">
      </template>
    </van-search>
  </div>
</template>
<script lang="ts">
import { defineComponent } from 'vue'
import { Search } from "vant";
import SearchIcon from "@/assets/icons/search-icon.svg"

export default defineComponent({
  setup() {
    return { SearchIcon };
  },
  components: {
    "van-search": Search,
  },
  props: {
    placeholder: {
      type: String,
      required: true,
    },
    modelValue: {
      type: String,
      required: true,
    },
  },
  emits: ['update:modelValue'],
  methods: {
    updateSearchQuery(searchQuery: string) {
      this.$emit('update:modelValue', searchQuery);
    }
  },
})
</script>
<style scoped lang="less">
.van-search {
  margin: 1.5em 0 1em 0;
  padding: 0;
  border-radius: 10em;
  box-shadow: -1px 1px 15px 15px rgba(103, 94, 226, 0.08);
}

#search-icon {
  padding-right: 1em;
  vertical-align: middle;
}

</style>
