<template>
<van-cell is-link @click="showPopup">Show Popup</van-cell>
<van-popup v-model="show">Content</van-popup>
<van-cell title="Select Single Date" :value="date" @click="show = true" />
<van-calendar v-model="show" @confirm="onConfirm" />
<Search placeholder="Search a friend" />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Popup, Search } from "vant"; 

export default defineComponent({
  components: {
    "van-popup": Popup,
    Search
  },
  data() {
    return {
      show: true,
      date: '',
    }
  },
});
</script>