<template>
  <div class="max-w-[90%] md:max-w-[80%] mx-auto">

    <div class="border-top-bottom text-h1 py-2 text-center md:text-left">
      <span>Food Menu</span>
    </div>

    <div class="text-h2 my-6 text-center md:text-left">
      <span class="font-bold">Plates</span>
    </div>

    <div class="grid grid-cols-1 justify-items-center md:grid-cols-2 lg:grid-cols-3 gap-4 mt-4">
      <menu-item v-for="item in plates" :imagePath="imagePath"  :key="item.id" :item="item" @expand-cart="expandCart" @add-to-cart="addToCart"></menu-item>
    </div>

    <div class="text-h2 my-6 text-center md:text-left">
      <span class="font-bold">Tacos</span>
    </div>

    <div class="grid grid-cols-1 justify-items-center md:grid-cols-2 lg:grid-cols-3 gap-4 mt-4">
      <menu-item v-for="item in tacos" :key="item.id" :item="item" @expand-cart="expandCart" @add-to-cart="addToCart"></menu-item>
    </div>

    <div class="text-h2 my-6 text-center md:text-left">
      <span class="font-bold">Tamales</span>
    </div>

    <div class="grid grid-cols-1 justify-items-center md:grid-cols-2 lg:grid-cols-3 gap-4 mt-4">
      <menu-item v-for="item in tamales" :key="item.id" :item="item" @expand-cart="expandCart" @add-to-cart="addToCart"></menu-item>
    </div>
  </div>
</template>

<script>
import MenuItem from './MenuItem.vue';
import ImagePath from '../assets/images/YH_May23_Chicken_Nachos.jpg';
import { globalState } from '../globalstate.js';

export default {
  components: {
    MenuItem
  },
  setup() {
    const cartItems = globalState.cartItems;

    const saveCart = () => {
      localStorage.setItem('cartItems', JSON.stringify(cartItems));
    };

    const addToCart = (item) => {
      cartItems.push(item);
      saveCart()
    };

    return {
      cartItems,
      addToCart
    };
  },
  data() {
    return {
      plates: [
        { id: 1, name: 'Carne Guisada Plate', price: 12.75, description: "Beef stew with poblano pepper and vegetables", imagePath: ImagePath },
        { id: 1, name: 'Asado de Puerco Plate', price: 12.75, description: "Pork stew with chile sauce and vegetables", imagePath: ImagePath },
        { id: 1, name: 'Bistek Ranchero Plate', price: 12.75, description: "Beef stew with poblano pepper and vegetables", imagePath: ImagePath },
      ],
      tacos: [
        { id: 3, name: 'Carne Guisada Tacos', price: 7.50, description: "Beef stew with poblano pepper and vegetables in a taco", imagePath: ImagePath },
        { id: 3, name: 'Asado de Puerco  Tacos', price: 7.50, description: "Beef stew with poblano pepper and vegetables in a taco", imagePath: ImagePath },
        { id: 3, name: 'Bistek Ranchero Tacos', price: 7.50, description: "Beef stew with poblano pepper and vegetables in a taco", imagePath: ImagePath },
      ],
      tamales: [
        { id: 2, name: 'Chicken Tamales', price: 15.75, description: "Authentic Chicken Mexican Tamale", imagePath: ImagePath },
        { id: 3, name: 'Pork Tamales', price: 7.50, description: "Authentic Pork Mexican Tamale", imagePath: ImagePath },
        { id: 1, name: 'Bean Tamales', price: 10.75, description: "Authentic Bean Mexican Tamale", imagePath: ImagePath },
        { id: 1, name: 'Cheese Tamales', price: 10.75 , description: "Authentic Cheese Mexican Tamale", imagePath: ImagePath},
      ]
    };
  },
  methods: {
    expandCart(){
      this.$emit('expand-cart');
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
