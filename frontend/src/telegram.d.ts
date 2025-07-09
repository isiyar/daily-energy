export {};

declare global {
  interface Window {
    Telegram: {
      WebApp: {
        initData: string;
        initDataUnsafe: any;
        close: () => void;
        sendData: (data: string) => void;
        expand: () => void;
        isExpanded: boolean;
        onEvent: (eventType: string, callback: () => void) => void;
        offEvent: (eventType: string, callback: () => void) => void;
        MainButton: {
          setText: (text: string) => void;
          show: () => void;
          hide: () => void;
          onClick: (callback: () => void) => void;
        };
        BackButton?: {
          show: () => void;
          hide: () => void;
          onClick: (callback: () => void) => void;
        };
      };
    };
  }
}
