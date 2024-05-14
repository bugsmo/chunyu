import { defineConfig, presetIcons, presetUno, presetAttributify } from 'unocss'
import UnocssIcons from '@unocss/preset-icons'

export default defineConfig({
  presets: [
    presetUno(),
    presetAttributify(),
    presetIcons({
      // 图标默认样式
      extraProperties: {
        display: 'inline-block',
        height: '1em',
        width: '1em',
      },
      /* options */
    }),
    UnocssIcons(),
  ],
})
