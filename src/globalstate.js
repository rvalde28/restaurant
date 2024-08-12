import { reactive } from 'vue';

// Initialize globalState
const globalState = reactive({
  cartItems: [],
  totalItems: 0,
});

function getTotalCount(){
  globalState.totalItems = 0
  const savedCartItems = localStorage.getItem('cartItems');
  if (savedCartItems) {
    globalState.cartItems = JSON.parse(savedCartItems);
  }

  for (let index = 0; index < globalState.cartItems.length; index++) {
    globalState.totalItems += globalState.cartItems[index].count;
  }
  console.log(" globalState.totalItems: ", globalState.totalItems)
}

function loadState() {
  const savedCartItems = localStorage.getItem('cartItems');
  if (savedCartItems) {
    globalState.cartItems = JSON.parse(savedCartItems);
  }
  getTotalCount()
}

loadState();

export { globalState, getTotalCount };