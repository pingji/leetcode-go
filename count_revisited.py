#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
统计README.md中Revisited列为Yes的行数
"""

import re

def count_revisited_yes():
    """统计README.md中Revisited列为Yes的行数"""
    
    # 读取README.md文件
    try:
        with open('README.md', 'r', encoding='utf-8') as f:
            content = f.read()
    except FileNotFoundError:
        print("错误：找不到README.md文件")
        return
    except Exception as e:
        print(f"错误：读取文件失败 - {e}")
        return
    
    # 使用正则表达式匹配表格行
    # 匹配格式：| 数字 | 标题 | 题名 | 难度 | 状态 | Yes |
    pattern = r'\|.*\|.*\|.*\|.*\|.*\|.*Yes.*\|'
    
    # 查找所有匹配的行
    matches = re.findall(pattern, content)
    
    # 统计结果
    total_revisited = len(matches)
    
    print(f"统计结果：")
    print(f"Revisited列为'Yes'的行数：{total_revisited}")
    
    # 提取题目信息并按编号排序
    problems = []
    for match in matches:
        parts = match.split('|')
        if len(parts) >= 3:
            number_str = parts[1].strip()
            title = parts[2].strip()
            
            # 提取题目编号（处理带前缀的情况如offer 04, interview 17.14等）
            if number_str.isdigit():
                number = int(number_str)
                prefix = ""
            else:
                # 处理带前缀的编号
                parts_num = number_str.split()
                if len(parts_num) >= 2:
                    prefix = parts_num[0]
                    try:
                        number = int(parts_num[1])
                    except ValueError:
                        number = 0
                else:
                    prefix = ""
                    number = 0
            
            problems.append((prefix, number, title))
    
    # 按编号排序
    problems.sort(key=lambda x: (x[0], x[1]))
    
    # 显示具体的题目信息
    if problems:
        print(f"\n按题目编号排序的题目列表：")
        for i, (prefix, number, title) in enumerate(problems, 1):
            if prefix:
                display_number = f"{prefix} {number}"
            else:
                display_number = str(number)
            print(f"{i:2d}. {display_number:>8} - {title}")
    
    return total_revisited

def count_by_category():
    """按分类统计Revisited为Yes的题目"""
    
    try:
        with open('README.md', 'r', encoding='utf-8') as f:
            content = f.read()
    except FileNotFoundError:
        print("错误：找不到README.md文件")
        return
    except Exception as e:
        print(f"错误：读取文件失败 - {e}")
        return
    
    # 按分类统计
    categories = {}
    current_category = ""
    
    lines = content.split('\n')
    for line in lines:
        # 检查是否是分类标题
        if line.startswith('## ') and not line.startswith('###'):
            current_category = line.replace('## ', '').strip()
            categories[current_category] = 0
        # 检查是否是表格行且Revisited为Yes
        elif '|' in line and 'Yes' in line and current_category:
            categories[current_category] += 1
    
    print(f"\n按分类统计：")
    total = 0
    for category, count in categories.items():
        if count > 0:
            print(f"{category}: {count}题")
            total += count
    
    print(f"\n总计：{total}题")
    return total

if __name__ == "__main__":
    print("开始统计README.md中Revisited为Yes的题目...")
    print("=" * 50)
    
    # 统计总数
    count_revisited_yes()
    
    print("\n" + "=" * 50)
    
    # 按分类统计
    count_by_category() 