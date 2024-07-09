import { createApp } from 'vue';
import App from './App.vue';
import { provideEventBus } from './EventBus'; // Adjust path as per your project structure

const app = createApp(App);

// Use provideEventBus to provide EventBus to the entire Vue app
app.use(provideEventBus);

app.mount('#app');