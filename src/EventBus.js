// EventBus.js

import { reactive, provide, inject } from 'vue';

const EventBusSymbol = Symbol('EventBus');

const useEventBus = () => {
  const events = reactive({});

  const emit = (event, ...args) => {
    if (!events[event]) return;
    events[event].forEach(callback => callback(...args));
  };

  const on = (event, callback) => {
    if (!events[event]) events[event] = [];
    events[event].push(callback);
  };

  return {
    emit,
    on,
  };
};

const provideEventBus = () => {
  const eventBus = useEventBus();
  provide(EventBusSymbol, eventBus);
};

const injectEventBus = () => {
  return inject(EventBusSymbol);
};

export { useEventBus, provideEventBus, injectEventBus };
