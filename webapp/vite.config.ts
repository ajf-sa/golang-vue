import reactRefresh from '@vitejs/plugin-react-refresh'


export default ({ command, mode }) => {
  if (command === 'serve') {
    return {
      plugins: [reactRefresh()],
      server:{port:5000,proxy:{
            '/api':{
                target: 'http://localhost:8888',
                changeOrigin: true
        
            },
            '/static':{
              target: 'http://localhost:8888',
              changeOrigin: true
      
          },
          },}
    }
  } else {
    return {
      build:{
        assetsDir:""
      },
      plugins: [reactRefresh()],
    }
  }
}