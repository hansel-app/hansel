<template>
<van-cell is-link @click="showPopup">Show Popup</van-cell>
<van-popup v-model="show">Content</van-popup>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Popup } from "vant"; 

export default defineComponent({
  components: {
    "van-popup": Popup,
  },
  data() {
    return {
      show: true,
      date: '',
    }
  },
  methods: {
    showPopup() {
      this.show = true;
    }
  }
});
</script>
