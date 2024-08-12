<template>
  
    <PopupWindow class="top-0 left-0 z-[1000]" :isOpen="isOpen" @close="isOpen = false">
      <!-- Popup content here -->
      <img class="p-0 rounded-t-lg" :src="item.imagePath">
      <h2 class="text-lg font-bold pt-4 pr-4 pl-4 pb-1">{{ item.name }}</h2>
      <p class="pt-1 pr-4 pl-4 pb-2">{{ item.description }}</p>

      <div class="rounded-b-lg overflow-hidden border-t box-cart-shadow">
        <!-- Header with buttons -->
        <div class="flex justify-between items-center bg-white px-4 py-2 border-b">
          <!-- Left buttons -->
          <div class="flex">
            <button @click="countSubtract()" class="p-2 rounded-full bg-blue-500 text-white hover:bg-blue-600">-</button>
            <div class="flex text-center w-12 pr-2 pl-2 bg-white rounded-none items-center justify-center"> {{ count  }} </div>
            <button @click="countAdd()" class="p-2 rounded-full bg-green-500 text-white hover:bg-green-600">+</button>
          </div>
          <!-- Right button -->
          <button @click="addToCart(item)"   class="flex items-center justify-around p-2 w-36 lg:w-48 rounded-full bg-gray-500 text-white hover:bg-gray-600">
              <div class=""> Add to Cart </div>
              <div class="">${{ count * item.price }}</div>
          </button>
        </div>

      </div>
    </PopupWindow>


    <button @click="isOpen = true" class="bg-white z-0 rounded-lg shadow-hover">
      <div class="rounded-lg">
        <img class="rounded-lg" src="../assets/images/YH_May23_Chicken_Nachos.jpg" ref="nachos">
      </div>
      <div class="p-3">
          <div class="text-menuCard text-left text-xl pb-2">
            {{ item.name }}
          </div>
          <div class="line"></div>
          <div class="text-menuCard text-left text-lg pt-8 pb-2">
              ${{ item.price }}
          </div >
      </div>
    </button>
</template>

<script>
import PopupWindow from './PopupWindow.vue';
import { globalState, getTotalCount } from '../globalstate.js';

const maxOrders = 25
export default {
  components: {
    PopupWindow,
  },
  props: ['item'],
  data() {
    return {
      count: 1,
      isOpen: false
    };
  },
  methods: {
    addToCart(item) {
      let itemInstance={
        name: item.name,
        price: item.price,
        count: this.count,
      }
      console.log("called")
      this.$emit('add-to-cart', itemInstance);
      this.$emit('expand-cart')
      this.count=1
      this.isOpen = false;
      getTotalCount();
    },
    countAdd(){
      if (this.count >= maxOrders){
        this.count = maxOrders
        return
      }
      this.count++
    },
    countSubtract(){
      if (this.count <= 1){
        this.count=1
        return
      } else{
        this.count--
      }
    }
  }
};
</script>
  
<style scoped>

.box-cart-shadow{
  box-shadow: 0 -2px 4px rgba(0, 0, 0, .11);
}

.line {
  bottom: 0;
  left: 50%;
  height: 1.75px;
  background-color: #DEC5E3;
}

.text-left{

}
.shadow-hover {
  transition: box-shadow 0.4s ease;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2); /* Default shadow */
}

.shadow-hover:hover {
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.6); /* Darker shadow that covers all sides */
}

/* Hide increment and decrement arrows on number input */
input[type="number"]::-webkit-outer-spin-button,
input[type="number"]::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
}


.item-count{
  font-size: .9rem;
  text-align: center;
}

input{
  border:0;
  text-align: center;
}

.center-text {
  margin: auto;
  width: 50%;
}

.center-div{
  display: flex; 
  flex-direction: row; 
  align-items: center; 
  justify-content: center;

}
.count-button{
  background-color: #57B8FF;
  color: black;
  border:0;
  padding:0;
  cursor:pointer;
}
.count-button:hover{
  background-color: #2a97e7;
}

.count-number{
  background-color: white;
  color: black
}
.item-price{
  margin-inline: auto;
  margin-bottom: 1vh;
}

.input-count-grid{
  display: grid;
  grid-template-columns: 4vw 10vw 4vw;
  background-color: white;
}
.cart-button-div{
  align-content: center;
  text-align: center;
}

.add-cart {
  align-content: center;
  text-align: center;
  padding:0.1vw 1vw;
  border-radius: 8px;
  color:#fff;
  background-color:#57B8FF;
  font-size:1.2vw;
  font-weight: 550 !important;
  border:0;
  cursor:pointer;
  margin:1em;
}

.item-name{
  color: white;
}

.item-price{
  color: white;
}

.box {
  background-color: #444;
  color: #fff;
}
img {
  position: relative;
  width: 100%;
  object-fit: contain;
  margin-left: auto;
  margin-right: auto;
}
</style>
  