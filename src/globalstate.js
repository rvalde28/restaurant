import { reactive } from 'vue';

// Initialize globalState
const globalState = reactive({
  cartItems: []
});

function loadState() {
  const savedCartItems = localStorage.getItem('cartItems');
  if (savedCartItems) {
    globalState.cartItems = JSON.parse(savedCartItems);
  }
}

loadState();

export { globalState };