<template>
  <div>

    <div class="wrapper border-top-bottom sm:max-w-screen-md xl:max-w-screen-xl text-h1 py-2">
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
import PopupWindow from './PopupWindow.vue';

export default {
  components: {
    MenuItem,
    PopupWindow
  },
  data() {
    return {
      plates: [
        { id: 1, name: 'Carne Guisada Plate', price: 12.75, description: "Beef stew with poblano pepper and vegetables" },
        { id: 1, name: 'Asado de Puerco Plate', price: 12.75, description: "Pork stew with chile sauce and vegetables" },
        { id: 1, name: 'Bistek Ranchero Plate', price: 12.75, description: "Beef stew with poblano pepper and vegetables" },
      ],
      tacos: [
        { id: 3, name: 'Carne Guisada Tacos', price: 7.50, description: "Beef stew with poblano pepper and vegetables in a taco" },
        { id: 3, name: 'Asado de Puerco  Tacos', price: 7.50, description: "Beef stew with poblano pepper and vegetables in a taco" },
        { id: 3, name: 'Bistek Ranchero Tacos', price: 7.50, description: "Beef stew with poblano pepper and vegetables in a taco" },
      ],
      tamales: [
        { id: 2, name: 'Chicken Tamales', price: 15.75, description: "Authentic Chicken Mexican Tamale" },
        { id: 3, name: 'Pork Tamales', price: 7.50, description: "Authentic Pork Mexican Tamale" },
        { id: 1, name: 'Bean Tamales', price: 10.75, description: "Authentic Bean Mexican Tamale" },
        { id: 1, name: 'Cheese Tamales', price: 10.75 , description: "Authentic Cheese Mexican Tamale"},
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
.border-top-bottom{
  border-top: 3px solid #57B8FF;
  border-bottom: 3px solid #57B8FF;  
}

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
    /* margin: 2rem; */
    margin-left: auto;
    margin-right: auto;
    /* grid-gap: 5vw; */
    /* background-color: #efece6; */
    color: #444;
}

/* .custom-gap-3{
  gap: 4.75rem;
} */

</style>
