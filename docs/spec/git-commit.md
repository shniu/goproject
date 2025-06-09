# Git Commit

## Commit Message 规范

Angular 规范的 Commit Message 格式如下：

```text
<Type>([optional scope]): <description>

[optional body]

[optional footer]
```

- Header 是必需的，Body 和 Footer 可以省略

### Type

- feat: [Production] 新增功能
- fix: [Production] 修复 Bug
- perf: [Production] 提高代码性能的变更
- refactor: [Production] 其他代码类的变更，但不属于 feat, fix, perf; 如简化代码，重命名，删除冗余代码等
- style: [Development] 代码格式类的变更，比如代码格式化、删除空行等
- test: [Development] 新增测试用例，或者更新现有测试用例
- ci: [Development] 持续集成或部署相关的改动，比如修改 Github Actions, Jenkins, 添加 CI 配置等，或更新 systemd unit 文件等
- docs: [Development] 文档更新，比如修改 README.md, 开发文档或用户文档等
- chore: [Development] 杂项，比如更新依赖，更新工具链，构建流程等

### Scope

scope 是用来说明 commit 的影响范围的，它必须是名词。如果 commit 影响范围较小，可以省略。

如果 commit 影响范围较大，可以考虑使用更具体的范围，比如某个模块或文件。

### Subject

subject 是 commit 的简短描述，必须以动词开头、使用现在时。比如，我们可以用
change，却不能用 changed 或 changes，而且这个动词的第一个字母必须是小写。通过
这个动词，我们可以明确地知道 commit 所执行的操作。此外我们还要注意，subject 的
结尾不能加英文句号。

## Body

Body 部分可以分成多行，而且格式也比较自由。不过，和 Header 里的一样，它也要以动
词开头，使用现在时。此外，它还必须要包括修改的动机，以及和跟上一版本相比的改动
点。

## Footer

Footer 部分不是必选的，可以根据需要来选择，主要用来说明本次 commit 导致的后果。
在实际应用中，Footer 通常用来说明不兼容的改动和关闭的 Issue 列表，
