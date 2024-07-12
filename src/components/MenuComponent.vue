<template>
  <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3 gap-4 wrapper max-w-60 sm:max-w-screen-md xl:max-w-screen-xl">
    <menu-item v-for="item in menuItems" :key="item.id" :item="item" @add-to-cart="addToCartHandler"></menu-item>
  </div>
</template>

<script>
import MenuItem from './MenuItem.vue';

export default {
  components: {
    MenuItem
  },
  data() {
    return {
      menuItems: [
        { id: 1, name: 'Carne Guisada', price: 10 },
        { id: 2, name: 'Tamales', price: 15 },
        { id: 3, name: 'Tacos', price: 7.5 },
        { id: 1, name: 'Carne Guisada', price: 10 },
        { id: 2, name: 'Tamales', price: 15 },
        { id: 3, name: 'Tacos', price: 7.5 },
      ]
    };
  },
  methods: {
    addToCartHandler(item) {
      this.$emit('add-to-cart', item);
    }
  }
};
</script>

<style scoped>
.center {
  margin: auto;
  width: 100%;
  border: 3px solid green;
  padding: 10px;
}

.center-div{
  max-width: fit-content;
  margin-left: auto;
  margin-right: auto;
}

.wrapper {
    position: relative;
    /* margin: 2rem; */
    margin-left: auto;
    margin-right: auto;
    display: grid;
    grid-gap: 5vw;
    /* background-color: #efece6; */
    color: #444;
}
</style>
