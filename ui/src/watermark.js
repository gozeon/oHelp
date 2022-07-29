class Watermask {
  constructor(opt) {
    this.opt = {
      el: document.body,
      cssText: 'position:absolute;left:0;top:0;display:block;z-index:-1;pointer-events:none;overflow:hidden;',
      text: 'hello world',
      textFont: '16px microsoft yahei',
      textColor: '#ededee',
      antiDeletion: true,
      resize: true,
      ...opt,
    }
    this.draw()
  }

  draw() {
    this.canvas && this.canvas.remove()

    const wrapperEl = this.opt.el

    const elHeight = wrapperEl?.clientHeight
    const elChildHeight = wrapperEl?.firstChild?.clientHeight

    const rect = {
      width: wrapperEl?.clientWidth,
      height: elHeight > elChildHeight ? elHeight : elChildHeight,
    }
    const canvas = document.createElement('canvas')
    canvas.width = rect.width
    canvas.height = rect.height
    canvas.style.cssText = this.opt.cssText
    wrapperEl.appendChild(canvas)
    this.canvas = canvas
    const xCount = rect.width / 20
    const yCount = rect.height / 50
    const ctx = canvas.getContext('2d')
    for (let i = 0; i < xCount; i++) {
      for (let j = 0; j < yCount; j++) {
        ctx.save()
        ctx.translate(i * 150 + 10, j * 100 + 60)
        ctx.rotate((-15 * Math.PI) / 180)
        ctx.fillStyle = this.opt.textColor
        ctx.font = this.opt.textFont

        if (i % 3 == 0) {
          ctx.fillText(this.opt.logoText || this.opt.text, 0, 0)
        } else if (i % 2 == 1) {
          var dateObj = new Date()
          var month = dateObj.getUTCMonth() + 1 //months from 1-12
          var day = dateObj.getUTCDate()
          var year = dateObj.getUTCFullYear()
          ctx.fillText(`${year}-${month}-${day}`, 0, 0)
        } else {
          ctx.fillText(this.opt.text, 0, 0)
        }

        ctx.restore()
      }
    }

    // Anti deletion 防删除
    this.opt.antiDeletion && this.ob()
    this.opt.resize && this.resize()
  }

  ob() {
    this.observer && this.observer.disconnect()
    this.observer = new MutationObserver((records) => {
      this.draw()
    })
    this.observer.observe(document.body, {
      attributes: true,
      childList: true,
      subtree: true,
    })
  }

  resize() {
    window.addEventListener('resize', () => {
      this.draw()
    })
  }

  destory() {
    this.observer && this.observer.disconnect()
    this.canvas && this.canvas.remove()

    window.removeEventListener('resize', () => {
      this.draw()
    })
  }
}

export { Watermask as default }
