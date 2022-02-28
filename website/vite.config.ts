import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';
import {quasar, transformAssetUrls} from '@quasar/vite-plugin';
import * as Path from 'path';
import viteCompression from 'vite-plugin-compression';
// https://vitejs.dev/config/
export default defineConfig({
  server: { host: '0.0.0.0' },
  build: {
    outDir: '../server/public'
  },
  resolve: {
    alias: {
      '@': Path.resolve(__dirname, 'src'),
    },
  },
  plugins: [
    vue({
      template: { transformAssetUrls },
    }),
    quasar({
      sassVariables: 'src/quasar-variables.sass',
    }),
    viteCompression(),
  ],
});
