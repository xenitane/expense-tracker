import { createGlobalStyle } from "styled-components";

const GlobalLayout = createGlobalStyle`
*{
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    list-style: none;
}

:root{
    --primary-color: #222260;
    --primary-color2: 'color: rgba(34,34,96,0.6)';
    --primary-color3: 'color: rgba(34,34,96,0.4)';
    --color-green: #42ad00;
    --color-grey:#aaaaaa;
    --color-accent: #f56692;
    --color-delete: #ff0000;
}
body{
    font-family: 'Nunito', sans-serif;
    font-size: clamp(1rem,1.5vw,1.2rem);
    overflow: hidden;
    color: rgb(34,34,96,0.6);
}
h1,h2,h3,h4,h5,h6{
    color:var(--primary-color);
}
`;

export default GlobalLayout;
