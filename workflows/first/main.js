const context = process.env.LAMRIM_CONTEXT
  ? JSON.parse(process.env.LAMRIM_CONTEXT)
  : {};

console.log(`workflow: ${process.env.LAMRIM_WORKFLOW}`);
console.log("variables:", context.variables);
