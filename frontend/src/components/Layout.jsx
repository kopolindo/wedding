import React from 'react';

const Layout = ({ children }) => {
  return (
    <div className="d-flex flex-column min-vh-100">
      <main className="flex-grow-1">{children}</main>
      <footer className="mt-auto Footer py-3 bg-light text-center">
      <div className="container">
        <span className="text-muted">
			Contatti:&emsp;Alex - 3294762906&emsp;Nadia - 3664818815
		</span>
		<br></br>
		<span className="text-muted small" style={{ fontStyle: 'italic' }}>
			Please be gentle, if you find something report to alex89.conti@gmail.com
      <br></br>
      Pull requests are well accepted https://github.com/kopolindo/wedding ![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
		</span>
      </div>
    </footer>
    </div>
  );
};

export default Layout;