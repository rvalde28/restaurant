<template>
  <div>

    <div class="wrapper sm:max-w-screen-md xl:max-w-screen-xl text-h1 py-4">
      <span >Food Menu</span>
    </div>

    <div class="wrapper sm:max-w-screen-md xl:max-w-screen-xl text-h2 pt-2 pb-8">
      <span >Plates</span>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3  wrapper gap-6 sm:gap-10 xl:gap-14 sm:max-w-screen-md xl:max-w-screen-xl">
      <menu-item v-for="item in plates" :key="item.id" :item="item" @add-to-cart="addToCartHandler"></menu-item>
    </div>

    <div class="wrapper sm:max-w-screen-md xl:max-w-screen-xl text-h2 pt-12 pb-8">
      <span >Tacos</span>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3  wrapper gap-6 sm:gap-10 xl:gap-14 max-w-60 sm:max-w-screen-md xl:max-w-screen-xl">
      <menu-item v-for="item in tacos" :key="item.id" :item="item" @add-to-cart="addToCartHandler"></menu-item>
    </div>

    <div class="wrapper sm:max-w-screen-md xl:max-w-screen-xl text-h2 pt-12 pb-8">
      <span >Tamales</span>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3  wrapper gap-6 sm:gap-10 xl:gap-14 max-w-60 sm:max-w-screen-md xl:max-w-screen-xl">
      <menu-item v-for="item in tamales" :key="item.id" :item="item" @add-to-cart="addToCartHandler"></menu-item>
    </div>
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
      plates: [
        { id: 1, name: 'Carne Guisada Plate', price: 12.75 },
        { id: 1, name: 'Asado de Puerco Plate', price: 12.75 },
        { id: 1, name: 'Bistek Ranchero Plate', price: 12.75 },
      ],
      tacos: [
        { id: 3, name: 'Carne Guisada Tacos', price: 7.50 },
        { id: 3, name: 'Asado de Puerco  Tacos', price: 7.50 },
        { id: 3, name: 'Bistek Ranchero Tacos', price: 7.50 },
      ],
      tamales: [
        { id: 2, name: 'Chicken Tamales', price: 15.75 },
        { id: 3, name: 'Pork Tamales', price: 7.5+0 },
        { id: 1, name: 'Bean Tamales', price: 10.75 },
        { id: 1, name: 'Cheese Tamales', price: 10.75 },
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
.text-h1{
  font-size: 2rem;
  font-weight: 450;
}

.text-h2{
  font-size: 1.75rem;
  font-weight: 450;
}

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
