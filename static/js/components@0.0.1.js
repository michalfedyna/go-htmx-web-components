import {css, html, LitElement} from '/static/lib/lit@3.0.0.js'

export class ExtraParagraphElement extends LitElement {
  static get styles() {
    return css`
      :host {
        color: red;
      }`;
  }

  static properties = {
    name: {type: String},
  }

  constructor() {
    super();
  }

  render() {
    console.log('render')
    return html`<p>Extra Paragraph</p>`;
  }
}

customElements.define('extra-p', ExtraParagraphElement);
