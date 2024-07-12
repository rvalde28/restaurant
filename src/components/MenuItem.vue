<template>
    <div class="bg-menuCard rounded-lg shadow-hover">
      <div class="rounded-lg">
        <img class="rounded-lg" src="../assets/images/YH_May23_Chicken_Nachos.jpg" ref="nachos">
      </div>
      <div class="p-3">
          <div class="text-white text-left text-xl">
            {{ item.name }}
          </div>

          <div class="text-white text-lg pt-10 pb-2">
              ${{ item.price }}
          </div >

          <!-- <div class="flex items-center justify-center py-2">
            <button class="px-4 bg-blue-500 text-white rounded-l" @click="countSubtract()">-</button>
            <input class="text-center w-24 px-2 bg-white border border-gray-300 rounded-none" type="number" v-model="count">
            <button class="px-4 bg-blue-500 text-white rounded-r" @click="countAdd()">+</button>
          </div> -->

          <!-- <div class="cart-button-div py-3">
            <button class="bg-menubtn hover:bg-blue-600 text-white font-bold py-1 px-4 rounded"  @click="addToCart(item)"> Add to Cart </button>
          </div > -->
      </div>
    </div>
</template>

<script>
  const maxOrders = 25
  export default {
    props: ['item'],
    data() {
    return {
      count: 1,
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
  