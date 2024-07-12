<template>
  <div class="grid grid-cols-1 sm:grid-cols-3 xl:grid-cols-4  wrapper gap-6 sm:gap-10 xl:gap-14 max-w-60 sm:max-w-screen-md xl:max-w-screen-xl">
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
        { id: 1, name: 'Carne Guisada Plate', price: 12.75 },
        { id: 1, name: 'Asado de Puerco Plate', price: 12.75 },
        { id: 1, name: 'Bistek Ranchero Plate', price: 12.75 },
        { id: 2, name: 'Chicken Tamales', price: 15.75 },
        { id: 3, name: 'Pork Tamales', price: 7.5+0 },
        { id: 1, name: 'Bean Tamales', price: 10.75 },
        { id: 1, name: 'Cheese Tamales', price: 10.75 },
        { id: 3, name: 'Carne Guisada Tacos', price: 7.50 },
        { id: 3, name: 'Asado de Puerco  Tacos', price: 7.50 },
        { id: 3, name: 'Bistek Ranchero Tacos', price: 7.50 },
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
    /* grid-gap: 5vw; */
    /* background-color: #efece6; */
    color: #444;
}

/* .custom-gap-3{
  gap: 4.75rem;
} */

</style>
