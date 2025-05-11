import {act} from 'react';
import ReactDOMClient from 'react-dom/client';
import { SubPage } from './page';


it(`can render and update a counter`,async () => {
  const container:HTMLElement = document.createElement('div');
  document.body.appendChild(container);

  await act(() => {
    ReactDOMClient.createRoot(container).render(<SubPage value={123} />);
  })
  
  expect(container.textContent).toContain('123');
})



it('future test', async () => {
  await act(async () => {
    // 将来の実装用
  });
});