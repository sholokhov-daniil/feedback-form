import typescript from '@rollup/plugin-typescript';
import terser from '@rollup/plugin-terser';

import './build/loadConfig.js';

export default {
  input: 'src/index.ts',   // точка входа
  output: {
    file: 'dist/sdk.js',   // результат
    format: 'umd',         // универсальный модуль для браузера и Node
    name: 'sform',           // имя глобальной переменной при использовании в браузере
    sourcemap: false,      // можно включить, если нужны карты
  },
  plugins: [
    typescript({ tsconfig: './tsconfig.json' }),
    terser()                // минификация
  ]
};