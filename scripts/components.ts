import { html, LitElement } from "lit";
import { customElement } from "lit/decorators.js";

@customElement("hello-world")
export class SimpleGreeting extends LitElement {
  render() {
    return html`<p>Hello World!</p>`;
  }
}
